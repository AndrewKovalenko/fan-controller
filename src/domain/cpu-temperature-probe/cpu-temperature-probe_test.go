package teperatureprobe

import (
	"testing"
)

const wantTemperatureCommandOutput = "temp=41.7'C"
const fakeMeasureTemperatureCommand = "echo " + wantTemperatureCommandOutput

func TestRunningCPUTempertatureCommand(t *testing.T) {
	commandOutput, err := runTemperatureCommand(fakeMeasureTemperatureCommand)

	if err != nil {
		t.Errorf("Error running command %s to test CPU temperature measurement. Error message: %s",
			fakeMeasureTemperatureCommand, err.Error())
	}

	if commandOutput != "" && commandOutput != wantTemperatureCommandOutput {
		t.Errorf("Invalid temperature measurement output: expected %s but recived %s",
			wantTemperatureCommandOutput, commandOutput)
	}

}
