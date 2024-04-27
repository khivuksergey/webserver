package webserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/khivuksergey/webserver/logger"
	"net/http"
	"time"
)

type server struct {
	httpServer   *http.Server
	router       http.Handler
	port         int
	stopTimeout  time.Duration
	stopHandlers []*StopHandler
	logger       logger.Logger
	srvCh        chan error
}

func NewServer(router http.Handler) Server {
	port, stopTimeout := initServerSettings(nil)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	srvCh := make(chan error, 1)

	return &server{
		httpServer:  httpServer,
		router:      router,
		port:        port,
		stopTimeout: stopTimeout,
		srvCh:       srvCh,
	}
}

func (s *server) WithConfig(config *ServerConfig) Server {
	port, stopTimeout := initServerSettings(config)
	s.httpServer.Addr = fmt.Sprintf(":%d", port)
	s.stopTimeout = stopTimeout
	return s
}

func (s *server) Start() (serverChannel chan error) {
	action, message, customMessage := "ServerStart", "Start", fmt.Sprintf("Starting server on port %d", s.port)
	s.logger.Info(logger.LogMessage{
		Action:        action,
		Message:       message,
		CustomMessage: &customMessage,
	})

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			action, message, customMessage = "ServerStart", "Start", fmt.Sprintf("Server start error: %v", err)
			s.logger.Error(logger.LogMessage{
				Action:        action,
				Message:       message,
				CustomMessage: &customMessage,
			})

			s.srvCh <- err
		}
	}()

	return s.srvCh
}

func (s *server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.stopTimeout)
	defer cancel()

	if len(s.stopHandlers) > 0 {
		defer func() {
			for _, handler := range s.stopHandlers {
				s.logger.Info(logger.LogMessage{
					Action:  "StopHandlers",
					Message: fmt.Sprintf("Stopping component %s", handler.Description),
				})

				err := handler.Stop()
				if err != nil {
					s.logger.Error(logger.LogMessage{
						Action:  "StopHandlers",
						Message: fmt.Sprintf("Error stopping component %s: %v", handler.Description, err),
					})
				}
			}
		}()
	}

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %v", err)
	}
	return nil
}

func (s *server) AddLogger(logger logger.Logger) Server {
	s.logger = logger
	return s
}

func (s *server) AddStopHandlers(stopHandlers ...*StopHandler) Server {
	s.stopHandlers = stopHandlers
	return s
}

func (s *server) IsLoggerMissing() bool {
	return s.logger == nil
}

func initServerSettings(config *ServerConfig) (port int, stopTimeout time.Duration) {
	port = DefaultPort
	stopTimeout = DefaultStopTimeout

	if config != nil {
		if config.Port > 0 {
			port = config.Port
		}
		if config.StopTimeout > time.Second {
			stopTimeout = config.StopTimeout
		}
	}
	return
}
