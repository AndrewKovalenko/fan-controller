package application

type LoggerInterface interface {
	Log(string)
	LogAndTerminate(string)
}
