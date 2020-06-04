package main

import app "fan-controller/src/application"

var controllerInitializer = app.InitializeFanController

const configFileName = "./fan-controller-config.yaml"

func main() {
	controllerInitializer(configFileName)
}
