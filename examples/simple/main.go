package main

import (
	"time"

	"mesh"
	"mesh/examples/simple/updateCounter"
	"mesh/pkg/features/configDirectory"
	"mesh/pkg/features/di"
	"mesh/pkg/features/logging"
)

func main() {
	app := mesh.New()

	app.AddFeature(&logging.Feature{
		LogLevel: -6,
	})

	app.AddFeature(&configDirectory.Feature{
		Parameters: configDirectory.Parameters{
			RootDirectory:     "~/.app-config",
			HotReloadInterval: time.Second,
		},
	})

	app.AddFeature(&di.Feature{})

	app.AddFeature(&updateCounter.Feature{})

	if err := app.Run(); err != nil {
		panic(err)
	}
}
