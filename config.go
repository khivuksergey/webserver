package webserver

import (
	"time"
)

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

type WebServerConfig struct {
	Port        int
	StopTimeout time.Duration
	Router      RouterConfig
}

type RouterConfig struct {
	UseLogger   bool
	UseRecovery bool
}

type SwaggerConfig struct {
	RoutePath   string
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}
