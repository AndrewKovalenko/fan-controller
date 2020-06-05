package application

import (
	"fmt"
	"time"

	controllerConfig "fan-controller/src/domain/controller-config"
	cpuTemperatureProbe "fan-controller/src/domain/cpu-temperature-probe"
	fanControl "fan-controller/src/domain/fan-control"
)

func InitializeFanController(controllerConfigFilePath string, logger LoggerInterface) error {
	defer func() {
		fanControl.CleanUp()
		logger.Log("Fan control CleanUp complete")
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

		fanSpeed := fanControllerConfig.GetFanSpeedSettingForTemperature(cpuTemperature)
		logMessage := fmt.Sprintf("CPU temperature: %f Fan Speed: %d", cpuTemperature, fanSpeed)
		logger.Log(logMessage)
		fanControl.SetFanSpeed(fanSpeed)

		time.Sleep(teperatureCheckingFrequency)
	}
}
