package mesh

import (
	"fmt"
	"sync"
	"time"

	ee "github.com/gravestench/eventemitter"
)

func New() Mesh {
	m := &mesh{}

	m.init()

	return m
}

type Mesh interface {
	Events() *ee.EventEmitter
	Features() []Feature
	AddFeature(feature Feature)
	RemoveFeature(feature Feature)
	Run() error
	Stop()
}

type mesh struct {
	bus          *ee.EventEmitter
	features     []Feature
	shuttingDown bool
}

func (m *mesh) init() {
	m.bus = ee.New()
	m.registerEvents()
}

func (m *mesh) Events() *ee.EventEmitter {
	return m.bus
}

func (m *mesh) Features() (copy []Feature) {
	return append(copy, m.features...)
}

func (m *mesh) hasFeature(feature Feature) bool {
	if m.indexOf(feature) != -1 {
		return true
	}

	return false
}

func (m *mesh) indexOf(feature Feature) int {
	for idx, other := range m.features {
		if feature == other {
			return idx
		}
	}

	return -1
}

func (m *mesh) AddFeature(feature Feature) {
	if m.hasFeature(feature) {
		return
	}

	m.registerEventsForFeature(feature)
	m.features = append(m.features, feature)
	m.Events().Emit(EventMeshFeatureAdded, feature)
}

func (m *mesh) RemoveFeature(feature Feature) {
	if !m.hasFeature(feature) {
		return
	}

	idx := m.indexOf(feature)
	m.features = append(m.features[:idx], m.features[idx+1:]...)
	m.Events().Emit(EventMeshFeatureRemoved, feature)
}

func (m *mesh) Run() error {
	var wg sync.WaitGroup

	for _, feature := range m.features {
		wg.Add(1)
		go func(f Feature) {
			defer wg.Done()
			m.initFeature(f)
		}(feature)
	}

	wg.Wait()

	lastUpdate := time.Now()

	for !m.shuttingDown {
		for _, feature := range m.features {
			wg.Add(1)

			go func(candidate Feature) {
				defer wg.Done()

				updater, ok := candidate.(UpdatingFeature)
				if !ok {
					return
				}

				updater.Update(time.Since(lastUpdate))
			}(feature)
		}

		wg.Wait() // wait for all updates to finish

		lastUpdate = time.Now()
	}

	return nil
}

func (m *mesh) initFeature(f Feature) {
	if readyAble, ok := f.(InitReadyFeature); ok {
		<-readyAble.Ready() // wait for feature to be ready
	}

	if err := f.Init(m); err != nil {
		m.Events().Emit(EventMeshFeatureInitError, fmt.Errorf("initializing %q: %v", f, err))
		return
	}

	m.Events().Emit(EventMeshFeatureInitialized, f)
}

func (m *mesh) Stop() {
	m.shuttingDown = true
}
