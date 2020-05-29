package teperatureprobe

import (
	"log"
	"os/exec"
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

	return string(commandOutput), err
}

func GetCPUTemperature() float32 {
	return 0
}
