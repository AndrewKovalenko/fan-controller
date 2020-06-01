package controllerconfig

import (
	"testing"
)

const testConfigFilePath = "../../cmd/fan-controller-config.yaml"

func TestReadYamlConfig(t *testing.T) {
	fanControllerConfig := ReadFanControllerConfig(testConfigFilePath)

	if len(fanControllerConfig.FanSpeedSettings) != 3 {
		t.Error("Controller fan speed config data is not parsed properly")
	}

	if fanControllerConfig.TemperatureUpdateFrequency != 20 {
		t.Error("Controller fan temperature update frequency is not parsed properly")
	}

	if fanControllerConfig.TurnOffTemperatureMargin != 5 {
		t.Error("Controller turn-off temerature margin is not parsed properly")
	}
}
