package configDirectory

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"mesh"
	"mesh/pkg/features/logging"
)

var _ mesh.Feature = &Feature{}

type Parameters struct {
	RootDirectory     string
	HotReloadInterval time.Duration
	HotReloadDisable  bool
}

type Feature struct {
	logging.LoggerIntegration
	Parameters
	updateWatchdog time.Duration
}

func (f *Feature) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &f.Parameters)
}

func (f *Feature) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(&f.Parameters, "", "\t")
}

func (f *Feature) Name() string {
	return "Config Directory Manager"
}

func (f *Feature) Init(mesh mesh.Mesh) (err error) {
	f.updateWatchdog = f.Parameters.HotReloadInterval

	if f.RootDirectory == "" { // handle default
		if f.RootDirectory, err = os.UserConfigDir(); err != nil {
			return fmt.Errorf("resolving user config dir: %v", err)
		}
	}

	return nil
}

func (f *Feature) Update(elapsed time.Duration) {
	f.updateWatchdog -= elapsed

	if f.updateWatchdog > 0 {
		return
	}

	f.updateWatchdog = f.Parameters.HotReloadInterval

	f.handleHotReload(elapsed)
}

func (f *Feature) handleHotReload(elapsed time.Duration) {
	if f.Parameters.HotReloadDisable {
		return
	}
}
