package teperatureprobe

import (
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
const one = 1

var logMessage = log.Output

func parseTemperature(commandOutputString string) (float32, error) {
	temperatureRegexp := regexp.MustCompile(cpuTemperatureResponsecPattern)
	matchedStrings := temperatureRegexp.FindStringSubmatch(commandOutputString)

	if len(matchedStrings) == one {
		temperatureValue := matchedStrings[0]
		cpuTemperature, err := strconv.ParseFloat(temperatureValue)

		if err != nil {
			logMessage(printToOutput, err.Error())
		}
		return cpuTemperature, err
	}

	return 0, error.New(fmt.Sprintf("Unable to parse CPU temperature form value %s", commandOutputString))
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
