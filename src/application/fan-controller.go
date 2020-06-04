package application

import (
	"time"

	controllerConfig "fan-controller/src/domain/controller-config"
	cpuTemperatureProbe "fan-controller/src/domain/cpu-temperature-probe"
	fanControl "fan-controller/src/domain/fan-control"
)

func getFanSpeed(temperature float32, fanSpeedSettings map[uint8]uint8) uint8 {

}

func InitializeFanController(controllerConfigFilePath string) error {
	defer func() {
		fanControl.CleanUp()
	}()

	fanControllerConfig, configReadingError :=
		controllerConfig.ReadFanControllerConfig(controllerConfigFilePath)

	if configReadingError != nil {
		return configReadingError
	}

	teperatureCheckingFrequency := time.Duration(fanControllerConfig.TemperatureUpdateFrequency) *
		time.Second
	fanControlInitError := fanControl.Init()

	if fanControlInitError != nil {
		return fanControlInitError
	}

	for {
		cpuTemperature, temperatureReadingError := cpuTemperatureProbe.GetCPUTemperature()

		if temperatureReadingError != nil {
			// log an error message
		}

		fanSpeed := getFanSpeed(cpuTemperature, fanControllerConfig.FanSpeedSettings)
		fanControl.SetFanSpeed(fanSpeed)

		time.Sleep(teperatureCheckingFrequency)
	}

	return nil
}
