package main

import (
	app "fan-controller/src/application"
	logger "fan-controller/src/domain/debugLogger"
	"os"
	"os/signal"
	"sync"
)

var wg = &sync.WaitGroup{}
var runFanController = app.RunFanController

const configFileName = "./fan-controller-config.yaml"

func main() {
	cancelChannel := make(chan os.Signal, 1)
	shutdownChannel := make(chan struct{})
	signal.Notify(cancelChannel, os.Interrupt)

	go func() {
		defer wg.Done()

		wg.Add(1)
		logger := logger.Logger{}
		runFanController(configFileName, logger, shutdownChannel)
	}()

	<-cancelChannel
	close(shutdownChannel)
	wg.Wait()
}
