package logger

import (
	"strings"
)

type LogLevel uint8

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Fatal
)

func GetLogLevelFromString(levelStr string) LogLevel {
	switch strings.ToUpper(levelStr) {
	case "DEBUG":
		return Debug
	case "INFO":
		return Info
	case "WARN":
		return Warn
	case "ERROR":
		return Error
	case "FATAL":
		return Fatal
	default:
		return Info
	}
}
