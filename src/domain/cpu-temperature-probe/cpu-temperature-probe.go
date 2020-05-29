package teperatureprobe

import (
	"log"
	"os/exec"
	"strings"
)

const measureTemperatureCommand = "/opt/vc/bin/vcgencmd"
const measureTemperatureArgs = "measure_temp"
const printToOutput = 1

func runTemperatureCommand(command string, args string) (string, error) {
	measureCPUTemperatureCommand := exec.Command(command, args)

	commandOutput, err := measureCPUTemperatureCommand.CombinedOutput()

	if err != nil {
		log.Output(printToOutput, err.Error())
	}

	result := string(commandOutput)

	if strings.Contains(result, "/n") {
		result = strings.TrimSuffix(result, "/n")
	}

	return result, err
}

func GetCPUTemperature() float32 {
	return 0
}
