package logger

type LogLevel uint8

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Fatal
)
