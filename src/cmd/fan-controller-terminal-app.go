package main

import app "fan-controller/src/application"

var controllerInitializer = app.InitializeFanController

func main() {
	controllerInitializer()
}
