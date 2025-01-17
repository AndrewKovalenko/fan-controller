package controllerconfig

import (
	"sort"
)

const maxUint8 = ^uint8(0)
const maxSpeed = 100
const turnOffFan = 0

type FanControllerConfig struct {
	TemperatureUpdateFrequency uint16 `yaml:"temperatureCheckFrequency"`
	TurnOffTemperatureMargin   uint8  `yaml:"stepdownTemperature"`

	FanSpeedSettings           map[uint8]uint8 `yaml:"fanSpeedSettings"`
	temperatureValuesAvailable []uint8
}

func getAllConfiguredTemperatureValues(fanConfig FanControllerConfig) []uint8 {
	result := make([]uint8, 0, len(fanConfig.FanSpeedSettings))

	for key := range fanConfig.FanSpeedSettings {
		result = append(result, key)
	}
	sort.Slice(result,
		func(i, j int) bool {
			return result[i] > result[j]
		})

	return result
}

func (f FanControllerConfig) GetFanSpeedSettingForTemperature(cpuTemperature float32,
	currentFanSpeed uint8) uint8 {
	roundedCPUTemperature := uint8(cpuTemperature)

	if f.temperatureValuesAvailable == nil {
		f.temperatureValuesAvailable = getAllConfiguredTemperatureValues(f)
	}

	for _, temperatureSetting := range f.temperatureValuesAvailable {
		if roundedCPUTemperature >= temperatureSetting {
			return f.FanSpeedSettings[temperatureSetting]
		}

		stepDownTemperature := temperatureSetting - f.TurnOffTemperatureMargin

		if roundedCPUTemperature >= stepDownTemperature &&
			currentFanSpeed == f.FanSpeedSettings[temperatureSetting] {
			return f.FanSpeedSettings[temperatureSetting]
		}
	}

	return turnOffFan
}
