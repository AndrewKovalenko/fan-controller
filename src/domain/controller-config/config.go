package controllerconfig

type FanControllerConfig struct {
	TemperatureUpdateFrequency int         `yaml:"temperatureCheckFrequency"`
	TurnOffTemperatureMargin   int         `yaml:"fanTurnOffTemperatureMargin"`
	FanSpeedSettings           map[int]int `yaml:"fanSpeedSettings"`
}
