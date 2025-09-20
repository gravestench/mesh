package logging

import (
	"log/slog"

	"mesh"
	"mesh/pkg/prettylog"
)

type Feature struct {
	LogLevel    int
	LogFile     string
	LogToStdout bool

	mesh mesh.Mesh
}

func (f *Feature) Name() string {
	return "Logging"
}

func (f *Feature) Init(mesh mesh.Mesh) error {
	f.mesh = mesh

	return nil
}

func (f *Feature) meshLogger() *slog.Logger {
	logger := slog.New(prettylog.NewHandler(&slog.HandlerOptions{
		Level: slog.Level(f.LogLevel),
	}))

	return logger
}

func (f *Feature) makeFeatureLogger(parent *Feature, feature mesh.Feature) *slog.Logger {
	logger := slog.New(prettylog.NewHandler(&slog.HandlerOptions{
		Level: slog.Level(parent.LogLevel),
	}))

	logger = logger.With("feature", feature.Name())

	return logger
}

func (f *Feature) OnMeshFeatureAdded(feature mesh.Feature) {
	if candidate, ok := feature.(loggingIntegration); ok {
		candidate.setLogger(f.makeFeatureLogger(f, feature))
		f.makeFeatureLogger(f, f).Debug("logger initialized", "target", feature.Name())
	}

	f.meshLogger().Debug(mesh.EventMeshFeatureAdded, "name", feature.Name())
}

func (f *Feature) OnMeshFeatureRemoved(feature mesh.Feature) {
	f.meshLogger().Debug(mesh.EventMeshFeatureRemoved, "name", feature.Name())
}

func (f *Feature) OnMeshFeatureInitError(err error) {
	f.meshLogger().Error(mesh.EventMeshFeatureInitError, "error", err)
}

func (f *Feature) OnMeshFeatureInitialized(feature mesh.Feature) {
	f.meshLogger().Debug(mesh.EventMeshFeatureInitialized, "name", feature.Name())
}
