package main

import (
	app "fan-controller/src/application"
	logger "fan-controller/src/domain/debugLogger"
)

var runFanController = app.RunFanController

const configFileName = "./fan-controller-config.yaml"

func main() {
	logger := logger.Logger{}
	runFanController(configFileName, logger)
}
