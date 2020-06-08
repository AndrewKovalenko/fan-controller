package main

import (
	"fan-controller/src/application"
	"testing"
)

var mockInitializerCalled = false

type applicationMock struct{}

func (app *applicationMock) InitializeFanController(string, application.LoggerInterface) error {
	mockInitializerCalled = true

	return nil
}

func TestMain(t *testing.T) {
	originalInitializer := runFanController
	appMock := new(applicationMock)
	runFanController = appMock.InitializeFanController

	main()

	if !mockInitializerCalled {
		t.Error("Terminal application doesn't call controller initialzer")
	}

	runFanController = originalInitializer
}
