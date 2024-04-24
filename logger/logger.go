package logger

type Logger interface {
	SetLevel(LogLevel)
	Debug(LogMessage)
	Info(LogMessage)
	Warn(LogMessage)
	Error(LogMessage)
	Fatal(LogMessage)
}

var Default = NewConsoleLogger()
