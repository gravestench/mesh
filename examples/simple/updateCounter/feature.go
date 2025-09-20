package updateCounter

import (
	"time"

	"mesh"
	"mesh/pkg/features/logging"
)

type Feature struct {
	logging.LoggerIntegration
	updates int
	dt      time.Duration
}

func (f *Feature) Name() string {
	return "Update Counter"
}

func (f *Feature) Init(mesh mesh.Mesh) error {
	return nil
}

func (f *Feature) Update(elapse time.Duration) {
	f.dt += elapse
	f.updates++

	if f.dt < time.Second {
		return
	}

	f.Log().Info("updates", "fps", f.updates)
	f.dt = 0
	f.updates = 0
}
