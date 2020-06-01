package controllerconfig

import "testing"

const testConfigFilePath = "../../cmd/fan-controller-config.yaml"

func TestReadYamlConfig(t *testing.T) {
	fanControllerConfig := ReadFanControllerConfig(testConfigFilePath)

	if fanControllerConfig.FanSpeedSettings == nil {
		t.Error("Controller fan speed config data can't be nil")
	}

}
