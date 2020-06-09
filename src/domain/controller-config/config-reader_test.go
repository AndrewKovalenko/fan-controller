package controllerconfig

import (
	"fmt"
	"strings"
	"testing"
)

const testConfigFilePath = "test-data/fan-controller-config.yaml"
const wrongFilePath = "test-data/not-exsist.yaml"
const notYamlFilePath = "test-data/text-file.txt"

const halfSpeed = 50
const twoThirdSpeed = 70

func TestReadYamlConfig(t *testing.T) {
	originalEvecutableDirectory := executableDirectory
	executableDirectory = ""
	fanControllerConfig, err := ReadFanControllerConfig(testConfigFilePath)

	if err != nil {
		t.Errorf("Config should be read successfully. Error: %s", err.Error())
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
	executableDirectory = originalEvecutableDirectory
}

func TestReadNotExistingConfig(t *testing.T) {
	originalEvecutableDirectory := executableDirectory
	executableDirectory = ""
	_, err := ReadFanControllerConfig(wrongFilePath)

	if err == nil {
		t.Error("Should throw an error trying to read not existing config")
	}

	expectedErrorMessage := fmt.Sprintf(fileNotFoundMessage, wrongFilePath)
	actualErrorMessage := err.Error()
	if actualErrorMessage != expectedErrorMessage {
		t.Errorf("Expected error message to be %s but got %s", expectedErrorMessage, actualErrorMessage)
	}
	executableDirectory = originalEvecutableDirectory
}

func TestReadInvalidYaml(t *testing.T) {
	originalEvecutableDirectory := executableDirectory
	executableDirectory = ""
	_, err := ReadFanControllerConfig(notYamlFilePath)

	if err == nil {
		t.Error("Should throw an error trying to read not existing config")
	}

	expectedErrorMessage := fmt.Sprintf(configParsingErrorMessage, notYamlFilePath, "")
	actualErrorMessage := err.Error()
	if !strings.HasPrefix(actualErrorMessage, expectedErrorMessage) {
		t.Errorf("Expect error message to start with %s but got %s", expectedErrorMessage, actualErrorMessage)
	}
	executableDirectory = originalEvecutableDirectory
}

func TestFillingInTemeratureScale(t *testing.T) {
	originalEvecutableDirectory := executableDirectory
	executableDirectory = ""
	fanControllerConfig, _ := ReadFanControllerConfig(testConfigFilePath)

	temperaturesConfigured := getAllConfiguredTemperatureValues(fanControllerConfig)

	if temperaturesConfigured == nil {
		t.Error("Temperature scale shouldn't be nil")
	}

	want := []uint8{60, 50, 40}
	for i, element := range want {
		if temperaturesConfigured[i] != element {
			t.Error("Temperatures available should be sorted")
		}
	}
	executableDirectory = originalEvecutableDirectory
}

func TestCalculatingFanSpeed(t *testing.T) {

	fanControllerConfig := FanControllerConfig{
		TemperatureUpdateFrequency: 5,
		TurnOffTemperatureMargin:   3,
		FanSpeedSettings: map[uint8]uint8{
			50: halfSpeed,
			54: twoThirdSpeed,
			58: maxSpeed},
	}

	fanSupposedToBeTurnedOff := fanControllerConfig.GetFanSpeedSettingForTemperature(49, 0)

	if fanSupposedToBeTurnedOff != turnOffFan {
		t.Error("Fan speed is not zero for temperature lower than lowest setting while wraming up")
	}

	fanSupposedToRunHalfSpeed := fanControllerConfig.GetFanSpeedSettingForTemperature(50, 0)

	if fanSupposedToRunHalfSpeed != halfSpeed {
		t.Error("Fan speed is not calsulated properly")
	}

	fanStillSupposedToRunHalfSpeed := fanControllerConfig.GetFanSpeedSettingForTemperature(49, halfSpeed)

	if fanStillSupposedToRunHalfSpeed != halfSpeed {
		t.Error("Fan supposed to maintain speed until cools down to stepDown temperature")
	}

	fanSupposedToRunTwoThirdSpeed := fanControllerConfig.GetFanSpeedSettingForTemperature(54, halfSpeed)

	if fanSupposedToRunTwoThirdSpeed != twoThirdSpeed {
		t.Error("Fan supposed to run two third of its speed")
	}
}
