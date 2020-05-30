package main

import app "../application"

var controllerInitializer controllerInitializerFunction = app.InitializeFanController

func main() {
	controllerInitializer()
}
