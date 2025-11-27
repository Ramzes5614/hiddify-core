package v2

import "fmt"

// CoreStartProbeFromConfig starts the core in probe mode using a fully built
// sing-box JSON configuration. It is designed for FFI/mobile callers and does
// not block the current goroutine.
func CoreStartProbeFromConfig(configJSON string) error {
	if configJSON == "" {
		return fmt.Errorf("config content is empty")
	}

	probeConfig := ConfigResult{
		Config:                configJSON,
		RefreshInterval:       0,
		HiddifyHiddifyOptions: nil,
	}

	return StartCoreWithConfig(probeConfig)
}

// CoreStopProbe stops a probe instance previously started via CoreStartProbeFromConfig.
func CoreStopProbe() error {
	return StopCore()
}
