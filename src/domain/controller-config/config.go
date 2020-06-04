package controllerconfig

import (
	"sort"
)

const maxUint8 = ^uint8(0)
const maxSpeed = 100
const turnOffFan = 0

type FanControllerConfig struct {
	TemperatureUpdateFrequency uint16 `yaml:"temperatureCheckFrequency"`
	TurnOffTemperatureMargin   uint8  `yaml:"fanTurnOffTemperatureMargin"`

	fanSpeedSettings           map[uint8]uint8 `yaml:"fanSpeedSettings"`
	temperatureValuesAvailable []uint8
}

func getAllConfiguredTemperatureValues(fanConfig FanControllerConfig) []uint8 {
	result := make([]uint8, len(fanConfig.fanSpeedSettings))

	for key, _ := range fanConfig.fanSpeedSettings {
		result = append(fanConfig.temperatureValuesAvailable, key)
	}

	return result
}

func (f FanControllerConfig) GetFanSpeedSettingForTemperature(cpuTemperature float32) uint8 {
	roundedCPUTemperature := uint8(cpuTemperature)

	if f.temperatureValuesAvailable == nil {
		f.temperatureValuesAvailable = getAllConfiguredTemperatureValues(f)
		sort.Slice(f.temperatureValuesAvailable,
			func(i, j int) bool {
				return f.temperatureValuesAvailable[i] < f.temperatureValuesAvailable[j]
			})
	}

	for index, temperatureSetting := range f.temperatureValuesAvailable {
		if roundedCPUTemperature > temperatureSetting {
			continue
		}

		stepDownTemperature := temperatureSetting - f.TurnOffTemperatureMargin

		if roundedCPUTemperature < stepDownTemperature && index == 0 {
			return turnOffFan
		}

		if roundedCPUTemperature < stepDownTemperature {
			lowerSpeedSettingIndex := uint8(index) - 1
			return f.fanSpeedSettings[lowerSpeedSettingIndex]
		}

		return f.fanSpeedSettings[temperatureSetting]
	}

	return maxSpeed
}
