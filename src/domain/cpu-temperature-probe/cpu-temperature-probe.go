package teperatureprobe

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

const measureTemperatureCommand = "/opt/vc/bin/vcgencmd"
const measureTemperatureArgs = "measure_temp"
const printToOutput = 1
const temperatureCaptureGroupName = "temperature"
const cpuTemperatureResponsecPattern = `.*=(?P<temperature>[\d\.]+)'C$`
const precission32 = 32

func findTemperatureString(expression *regexp.Regexp, matches []string) (string, bool) {
	groupNmaes := expression.SubexpNames()

	for i, matchString := range matches {
		if groupNmaes[i] == temperatureCaptureGroupName {
			return matchString, true
		}
	}

	return "", false
}

func parseTemperature(commandOutputString string) (float32, error) {
	temperatureRegexp := regexp.MustCompile(cpuTemperatureResponsecPattern)
	matchedStrings := temperatureRegexp.FindStringSubmatch(commandOutputString)

	temperatureString, temperatureValueFound := findTemperatureString(temperatureRegexp, matchedStrings)

	if temperatureValueFound {
		cpuTemperature, err := strconv.ParseFloat(temperatureString, precission32)

		if err != nil {
			return 0, err
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
		return "", err
	}

	return string(commandOutput), err
}

func GetCPUTemperature() (float32, error) {
	temperatureString, readCPUTemperatureError := runTemperatureCommand(measureTemperatureCommand,
		measureTemperatureArgs)

	if readCPUTemperatureError != nil {
		return 0, readCPUTemperatureError
	}

	cpuTemperature, parsingTemperatureError := parseTemperature(temperatureString)

	return cpuTemperature, parsingTemperatureError
}
