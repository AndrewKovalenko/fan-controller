package controllerconfig

import (
	"errors"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

const fileNotFoundMessage = "Unable to find config file %s"
const configParsingErrorMessage = "Unable to parse yaml config: %s \n Error: %s"

func ReadFanControllerConfig(filename string) (FanControllerConfig, error) {
	configContent, err := ioutil.ReadFile(filename)

	if err != nil {
		errorMessge := fmt.Sprintf(fileNotFoundMessage, filename)
		return FanControllerConfig{}, errors.New(errorMessge)
	}

	fanControllerConfig := FanControllerConfig{}
	parsingConfigError := yaml.Unmarshal(configContent, &fanControllerConfig)

	if parsingConfigError != nil {
		parsingErrorMessage := fmt.Sprintf(configParsingErrorMessage, string(filename), parsingConfigError)
		return FanControllerConfig{}, errors.New(parsingErrorMessage)
	}

	return fanControllerConfig, nil
}
