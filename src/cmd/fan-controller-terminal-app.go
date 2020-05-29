package main

import app "../application"

type controllerInitializerFunction func()

var controllerInitializer controllerInitializerFunction = app.InitializeFanController

func main() {
	controllerInitializer()
}
