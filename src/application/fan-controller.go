package application

import (
	"fmt"
	"time"

	controllerConfig "fan-controller/src/domain/controller-config"
	cpuTemperatureProbe "fan-controller/src/domain/cpu-temperature-probe"
	fanControl "fan-controller/src/domain/fan-control"
)

func RunFanController(controllerConfigFilePath string, logger LoggerInterface) error {
	defer func() {
		fanControl.CleanUp()
		logger.Log("Fan control CleanUp complete")
	}()

	fanControllerConfig, configReadingError :=
		controllerConfig.ReadFanControllerConfig(controllerConfigFilePath)

	if configReadingError != nil {
		logMessage := fmt.Sprintf("config reading error %s", configReadingError.Error())
		logger.Log(logMessage)
		return configReadingError
	}

	logger.Log("Config read sucessfully")

	teperatureCheckingFrequency := time.Duration(fanControllerConfig.TemperatureUpdateFrequency) *
		time.Second
	fanControlInitError := fanControl.Init()

	if fanControlInitError != nil {
		logMessage := fmt.Sprintf("Fan control init error %s", fanControlInitError.Error())
		logger.Log(logMessage)
		return fanControlInitError
	}

	for {
		cpuTemperature, temperatureReadingError := cpuTemperatureProbe.GetCPUTemperature()

		if temperatureReadingError != nil {
			logMessage := fmt.Sprintf("Temperature reading error %s", temperatureReadingError.Error())
			logger.Log(logMessage)
		}

		fanSpeed := fanControllerConfig.GetFanSpeedSettingForTemperature(cpuTemperature)
		fanControl.SetFanSpeed(fanSpeed)

		time.Sleep(teperatureCheckingFrequency)
	}
}
