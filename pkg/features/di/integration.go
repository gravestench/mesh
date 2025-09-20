package di

import (
	"mesh"
)

type Resolver interface {
	ResolveDependencies(features ...mesh.Feature) (success bool)
}
