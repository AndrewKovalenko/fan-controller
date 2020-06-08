package controllerconfig

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

const fileNotFoundMessage = "Unable to find config file %s"
const configParsingErrorMessage = "Unable to parse yaml config: %s \n Error: %s"

func ReadFanControllerConfig(filename string) (FanControllerConfig, error) {
	executableDirectory, directoryReadingError := filepath.Abs(filepath.Dir(os.Args[0]))
	if directoryReadingError != nil {
		return FanControllerConfig{}, directoryReadingError
	}

	configFilePath := path.Join(executableDirectory, filename)
	configContent, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		errorMessge := fmt.Sprintf(fileNotFoundMessage, configFilePath)
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
