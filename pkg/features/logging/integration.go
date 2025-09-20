package logging

import (
	"log/slog"
)

// LoggerIntegration can be embedded into another feature
// to automatically get a logger assigned
type LoggerIntegration struct {
	*slog.Logger
	ready chan bool
}

// this is a private interface satisfied by the above struct. this forces
// the user to use the above struct when embedding in another feature.
type loggingIntegration interface {
	Log() *slog.Logger
	setLogger(logger *slog.Logger)
	loggingIntegration()
}

func (li *LoggerIntegration) Log() *slog.Logger {
	for li.Logger != nil {
		return li.Logger
	}

	for li.ready == nil {
	}

	<-li.ready

	return li.Logger
}

func (li *LoggerIntegration) setLogger(logger *slog.Logger) {
	if li.ready == nil {
		li.ready = make(chan bool, 1)
	}

	li.Logger = logger
	li.ready <- true
}

// only we implement this
func (*LoggerIntegration) loggingIntegration() {
	return
}
