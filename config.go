package webserver

import (
	"time"
)

const (
	DefaultPort        = 8080
	DefaultStopTimeout = 5 * time.Second
)

var DefaultServerConfig = ServerConfig{
	Port:        DefaultPort,
	StopTimeout: DefaultStopTimeout,
}

var DefaultRouterConfig = RouterConfig{
	UseLogger:   false,
	UseRecovery: true,
}

type ServerConfig struct {
	Port        int
	StopTimeout time.Duration
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
