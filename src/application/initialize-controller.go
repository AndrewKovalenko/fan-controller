package application

import (
	"fmt"

	controllerConfig "fan-controller/src/domain/controller-config"
)

func InitializeFanController(controllerConfigFilePath string) error {
	fanControllerConfig, configReadingError :=
		controllerConfig.ReadFanControllerConfig(controllerConfigFilePath)

	if configReadingError != nil {
		return configReadingError
	}

	fmt.Println(fanControllerConfig.TemperatureUpdateFrequency)

	return nil
}
