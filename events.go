package mesh

const (
	EventMeshFeatureAdded       = "feature added"
	EventMeshFeatureRemoved     = "feature removed"
	EventMeshFeatureInitError   = "feature init error"
	EventMeshFeatureInitialized = "feature init"
)

func (m *mesh) registerEvents() {
	m.Events().On(EventMeshFeatureAdded, func(args ...any) {
		if len(args) < 1 {
			return
		}

		if feature, ok := args[0].(Feature); ok {
			m.AddFeature(feature)
		}
	})

	m.Events().On(EventMeshFeatureRemoved, func(args ...any) {
		if len(args) < 1 {
			return
		}

		if feature, ok := args[0].(Feature); ok {
			m.RemoveFeature(feature)
		}
	})
}

func (m *mesh) registerEventsForFeature(feature Feature) {
	if candidate, ok := feature.(EventHandlerMeshFeatureAdded); ok {
		m.Events().On(EventMeshFeatureAdded, func(args ...any) {
			candidate.OnMeshFeatureAdded(args[0].(Feature))
		})
	}

	if candidate, ok := feature.(EventHandlerMeshFeatureRemoved); ok {
		m.Events().On(EventMeshFeatureRemoved, func(args ...any) {
			candidate.OnMeshFeatureRemoved(args[0].(Feature))
		})
	}

	if candidate, ok := feature.(EventHandlerMeshFeatureInitError); ok {
		m.Events().On(EventMeshFeatureInitError, func(args ...any) {
			candidate.OnMeshFeatureInitError(args[0].(error))
		})
	}

	if candidate, ok := feature.(EventHandlerMeshFeatureInitialized); ok {
		m.Events().On(EventMeshFeatureInitialized, func(args ...any) {
			candidate.OnMeshFeatureInitialized(args[0].(Feature))
		})
	}

}

type EventHandlerMeshFeatureAdded interface {
	OnMeshFeatureAdded(feature Feature) // arg is the feature that was added
}

type EventHandlerMeshFeatureRemoved interface {
	OnMeshFeatureRemoved(feature Feature) // arg is the feature that was removed
}

type EventHandlerMeshFeatureInitError interface {
	OnMeshFeatureInitError(err error) // err is wrapped and shows feature name
}

type EventHandlerMeshFeatureInitialized interface {
	OnMeshFeatureInitialized(feature Feature) // arg is the feature that init'd
}
