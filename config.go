package webserver

import "time"

const (
	DefaultPort        int           = 8080
	DefaultStopTimeout time.Duration = 5 * time.Second
)

var DefaultWebServerConfig = WebServerConfig{
	Port:        DefaultPort,
	StopTimeout: DefaultStopTimeout,
	Router: RouterConfig{
		UseLogger:   true,
		UseRecovery: true,
	},
}

type Configuration struct {
	WebServer WebServerConfig
}

type WebServerConfig struct {
	Port        int
	StopTimeout time.Duration
	Router      RouterConfig
}

type RouterConfig struct {
	UseLogger   bool
	UseRecovery bool
}
