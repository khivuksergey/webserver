package logger

import "fmt"

type logger struct {
	level LogLevel
}

func NewConsoleLogger() Logger {
	return &logger{
		level: Info,
	}
}

func (l *logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *logger) Debug(message LogMessage) {
	if l.level > Debug {
		return
	}
	fmt.Printf(message.String(), "DEBUG")
}

func (l *logger) Info(message LogMessage) {
	if l.level > Info {
		return
	}
	fmt.Printf(message.String(), "INFO")
}

func (l *logger) Warn(message LogMessage) {
	if l.level > Warn {
		return
	}
	fmt.Printf(message.String(), "WARN")
}

func (l *logger) Error(message LogMessage) {
	if l.level > Error {
		return
	}
	fmt.Printf(message.String(), "ERROR")
}

func (l *logger) Fatal(message LogMessage) {
	fmt.Printf(message.String(), "FATAL")
}
