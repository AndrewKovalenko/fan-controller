package filereader

import (
	"io/ioutil"
	"log"
)

const printToOutput = 1

func ReadFile(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Output(printToOutput, err.Error())
	}

	return content, err
}
