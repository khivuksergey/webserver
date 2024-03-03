package logger

import "fmt"

type logger struct{}

func NewConsoleLogger() Logger {
	return &logger{}
}

func (l logger) Debug(message LogMessage) {
	fmt.Printf(message.String(), "DEBUG")
}

func (l logger) Info(message LogMessage) {
	fmt.Printf(message.String(), "INFO")
}

func (l logger) Warn(message LogMessage) {
	fmt.Printf(message.String(), "WARN")
}

func (l logger) Error(message LogMessage) {
	fmt.Printf(message.String(), "ERROR")
}

func (l logger) Fatal(message LogMessage) {
	fmt.Printf(message.String(), "FATAL")
}
