package controllerconfig

type FanControllerConfig struct {
	TemperatureUpdateFrequency uint16          `yaml: "temperatureCheckFrequency"`
	TurnOffTemperatureMargin   uint8           `yaml: "fanTurnOffTemperatureMargin"`
	FanSpeedSettings           map[uint8]uint8 `yaml: "fanSpeedSettings"`
}
