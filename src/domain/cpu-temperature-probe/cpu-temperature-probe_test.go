package teperatureprobe

import (
	"strings"
	"testing"
)

func TestRunningCPUTempertatureCommand(t *testing.T) {
	const wantTemperatureCommandOutput = "temp=41.7'C"
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
	messageLoged := false
	originalLogger := logMessage

	mockLogger := func(int, string) error {
		messageLoged = true
		return nil
	}

	logMessage = mockLogger

	runTemperatureCommand("Hello", "there")

	if !messageLoged {
		t.Error("CPU temperature measurement function doesn't log error message when command fails")
	}

	logMessage = originalLogger
}
