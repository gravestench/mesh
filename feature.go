package mesh

import (
	"time"
)

type Feature interface {
	Name() string
	Init(mesh Mesh) error
}

type InitReadyFeature interface {
	Feature
	Ready() chan bool
}

type UpdatingFeature interface {
	Feature
	Update(elapse time.Duration)
}
