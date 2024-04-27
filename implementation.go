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
	stopHandlers *[]StopHandler
	logger       logger.Logger
	srvCh        chan error
}

func NewServer(wsConfig *WebServerConfig, router http.Handler) Server {
	port, stopTimeout := initServerSettings(wsConfig)

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

	if s.stopHandlers != nil {
		defer func() {
			for _, handler := range *s.stopHandlers {
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

func (s *server) AddLogger(logger logger.Logger) {
	s.logger = logger
}

func (s *server) AddStopHandlers(stopHandlers *[]StopHandler) {
	s.stopHandlers = stopHandlers
}

func (s *server) IsLoggerMissing() bool {
	return s.logger == nil
}

func initServerSettings(wsConfig *WebServerConfig) (port int, stopTimeout time.Duration) {
	port = DefaultPort
	stopTimeout = DefaultStopTimeout

	if wsConfig != nil {
		if wsConfig.Port > 0 {
			port = wsConfig.Port
		}
		if wsConfig.StopTimeout > time.Second {
			stopTimeout = wsConfig.StopTimeout
		}
	}
	return
}
