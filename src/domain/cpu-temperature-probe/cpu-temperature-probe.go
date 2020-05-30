package teperatureprobe

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

const measureTemperatureCommand = "/opt/vc/bin/vcgencmd"
const measureTemperatureArgs = "measure_temp"
const printToOutput = 1
const cpuTemperatureResponsecPattern = `.*=([\d\.]+)'C$`
const precission32 = 32

var logMessage = log.Output

func hasOnlyOneElement(collection []string) bool {
	return len(collection) == 1
}

func parseTemperature(commandOutputString string) (float32, error) {
	temperatureRegexp := regexp.MustCompile(cpuTemperatureResponsecPattern)
	matchedStrings := temperatureRegexp.FindStringSubmatch(commandOutputString)

	if hasOnlyOneElement(matchedStrings) {
		temperatureValue := matchedStrings[0]
		cpuTemperature, err := strconv.ParseFloat(temperatureValue, precission32)

		if err != nil {
			logMessage(printToOutput, err.Error())
		}
		return float32(cpuTemperature), err
	}

	wrongCommandOutputMessage := fmt.Sprintf("Unable to parse CPU temperature form value %s", commandOutputString)
	return 0, errors.New(wrongCommandOutputMessage)
}

func runTemperatureCommand(command string, args string) (string, error) {
	measureCPUTemperatureCommand := exec.Command(command, args)

	commandOutput, err := measureCPUTemperatureCommand.CombinedOutput()

	if err != nil {
		logMessage(printToOutput, err.Error())
	}

	return string(commandOutput), err
}

func GetCPUTemperature() float32 {
	return 0
}
