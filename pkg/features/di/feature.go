package di

import (
	"time"

	"mesh"
)

type Feature struct {
	mesh mesh.Mesh
}

func (f *Feature) Name() string {
	return "Dependency Injection"
}

func (f *Feature) Init(mesh mesh.Mesh) error {
	f.mesh = mesh

	return nil
}

func (f *Feature) Update(elapse time.Duration) {
	for f.mesh == nil {
		// dwell
	}

	for _, feature := range f.mesh.Features() {
		f.resolveDependenciesForFeature(feature)
	}
}

func (f *Feature) OnMeshFeatureAdded(feature mesh.Feature) {
	f.resolveDependenciesForFeature(feature)
}

func (f *Feature) resolveDependenciesForFeature(feature mesh.Feature) bool {
	for f.mesh == nil {
		// dwell
	}

	resolver, ok := feature.(Resolver)
	if !ok {
		return true
	}

	return resolver.ResolveDependencies(f.mesh.Features()...)
}
