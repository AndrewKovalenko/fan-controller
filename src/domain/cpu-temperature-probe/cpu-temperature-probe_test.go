package teperatureprobe

import (
	"math"
	"strings"
	"testing"
)

const wantTemperatureCommandOutput = "temp=41.7'C"
const wantTemperatureValue = 41.7

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float32) bool {
	return math.Abs(float64(a-b)) <= float64EqualityThreshold
}

func TestRunningCPUTempertatureCommand(t *testing.T) {
	const fakeMeasureTemperatureCommand = "echo"

	commandOutput, err := runTemperatureCommand(fakeMeasureTemperatureCommand,
		wantTemperatureCommandOutput)

	if err != nil {
		t.Errorf("Error running command %s to test CPU temperature measurement. Error message: %s",
			fakeMeasureTemperatureCommand, err.Error())
	}

	if commandOutput != "" && !strings.Contains(commandOutput, wantTemperatureCommandOutput) {
		t.Errorf("Invalid temperature measurement output: expected %s but recived %s",
			wantTemperatureCommandOutput, commandOutput)
	}
}

func TestRunningCPUTemperatureCommandLogsMessageIfError(t *testing.T) {
	_, err := runTemperatureCommand("Hello", "there")

	if err == nil {
		t.Error("CPU temperature measurement function doesn't log error message when command fails")
	}
}

func TestParseTemperature(t *testing.T) {
	temperatureValue, err := parseTemperature(wantTemperatureCommandOutput)

	if err != nil {
		t.Errorf("Error parsing valid temperature value %s. Error message: %s", wantTemperatureCommandOutput, err)
	}

	if err == nil && !almostEqual(temperatureValue, wantTemperatureValue) {
		t.Errorf("Expected temperature value to be %f but recived %f", wantTemperatureValue, temperatureValue)
	}
}
