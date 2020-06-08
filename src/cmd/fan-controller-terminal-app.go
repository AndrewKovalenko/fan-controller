package main

import (
	app "fan-controller/src/application"
	logger "fan-controller/src/domain/debugLogger"
	"os"
	"os/signal"
	"sync"
)

var shutdownChanel = make(chan struct{})
var wg = &sync.WaitGroup{}
var runFanController = app.RunFanController

const configFileName = "./fan-controller-config.yaml"

func main() {
	go func() {
		defer wg.Done()

		logger := logger.Logger{}
		runFanController(configFileName, logger, shutdownChanel)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	close(shutdownChanel)
	wg.Wait()
}
