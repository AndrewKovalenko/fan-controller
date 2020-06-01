package controllerconfig

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v3"
)

func ReadFanControllerConfig(filename string) FanControllerConfig {
	configContent, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("Unable to find config file %s", filename)
	}

	fanControllerConfig := FanControllerConfig{}
	parsingConfigError := yaml.Unmarshal(configContent, &fanControllerConfig)

	if parsingConfigError != nil {
		log.Fatalf("Unable to parse yaml config: %s \n Error: %s",
			string(configContent), parsingConfigError)
	}

	return fanControllerConfig
}
