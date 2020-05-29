package main

import "testing"

var mockInitializerCalled = false

type applicationMock struct{}

func (app *applicationMock) InitializeFanController() {
	mockInitializerCalled = true
}

func TestMain(t *testing.T) {
	originalInitializer := controllerInitializer
	appMock := new(applicationMock)
	controllerInitializer = appMock.InitializeFanController

	main()

	if !mockInitializerCalled {
		t.Error("Terminal application doesn't call controller initialzer")
	}

	controllerInitializer = originalInitializer
}
