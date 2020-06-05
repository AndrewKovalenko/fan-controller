package controllerconfig

import (
	"fmt"
	"strings"
	"testing"
)

const testConfigFilePath = "./test-data/fan-controller-config.yaml"
const wrongFilePath = "./test-data/not-exsist.yaml"
const notYamlFilePath = "./test-data/text-file.txt"

func TestReadYamlConfig(t *testing.T) {
	fanControllerConfig, err := ReadFanControllerConfig(testConfigFilePath)

	if err != nil {
		t.Error("Config should be read successfully")
	}

	if len(fanControllerConfig.FanSpeedSettings) != 3 {
		t.Error("Controller fan speed config data is not parsed properly")
	}

	if fanControllerConfig.TemperatureUpdateFrequency != 20 {
		t.Error("Controller fan temperature update frequency is not parsed properly")
	}

	if fanControllerConfig.TurnOffTemperatureMargin != 5 {
		t.Error("Controller turn-off temerature margin is not parsed properly")
	}
}

func TestReadNotExistingConfig(t *testing.T) {
	_, err := ReadFanControllerConfig(wrongFilePath)

	if err == nil {
		t.Error("Should throw an error trying to read not existing config")
	}

	expectedErrorMessage := fmt.Sprintf(fileNotFoundMessage, wrongFilePath)
	actualErrorMessage := err.Error()
	if actualErrorMessage != expectedErrorMessage {
		t.Errorf("Expected error message to be %s but got %s", expectedErrorMessage, actualErrorMessage)
	}
}

func TestReadInvalidYaml(t *testing.T) {
	_, err := ReadFanControllerConfig(notYamlFilePath)

	if err == nil {
		t.Error("Should throw an error trying to read not existing config")
	}

	expectedErrorMessage := fmt.Sprintf(configParsingErrorMessage, notYamlFilePath, "")
	actualErrorMessage := err.Error()
	if !strings.HasPrefix(actualErrorMessage, expectedErrorMessage) {
		t.Errorf("Expect error message to start with %s but got %s", expectedErrorMessage, actualErrorMessage)
	}
}

func TestFillingInTemeratureScale(t *testing.T) {
	fanControllerConfig, _ := ReadFanControllerConfig(testConfigFilePath)

	temperaturesConfigured := getAllConfiguredTemperatureValues(fanControllerConfig)

	if temperaturesConfigured == nil {
		t.Error("Temperature scale shouldn't be nil")
	}

	want := []uint8{40, 50, 60}
	for i, element := range want {
		if temperaturesConfigured[i] != element {
			t.Error("Temperatures available should be sorted")
		}
	}
}
