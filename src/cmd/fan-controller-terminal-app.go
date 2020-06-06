package main

import (
	app "fan-controller/src/application"
	logger "fan-controller/src/domain/debugLogger"
)

var controllerInitializer = app.InitializeFanController

const configFileName = "./fan-controller-config.yaml"

func main() {
	logger := logger.Logger{}
	controllerInitializer(configFileName, logger)
}
