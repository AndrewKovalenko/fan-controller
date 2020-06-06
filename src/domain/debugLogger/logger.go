package debugLogger

import "log"

type Logger struct{}

func (l Logger) Log(message string) {
	log.Println(message)
}

func (l Logger) LogAndTerminate(message string) {
	log.Fatalln(message)
}
