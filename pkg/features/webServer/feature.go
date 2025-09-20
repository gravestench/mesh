package webServer

import (
	"time"

	"mesh"
	"mesh/pkg/features/logging"
)

type Feature struct {
	logging.LoggerIntegration
}

func (f *Feature) Name() string {
	return "Web Server"
}

func (f *Feature) Init(c mesh.Mesh) error {

	return nil
}

func (f *Feature) Update(elapse time.Duration) {

}
