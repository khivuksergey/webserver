package webserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/khivuksergey/webserver/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	Start() chan error
	Stop() error
	AddLogger(logger.Logger)
	isLoggerMissing() bool
}

type ServerOptions struct {
	//StopTimeoutMS int
	UseLogger bool
}

type server struct {
	router        http.Handler
	httpServer    *http.Server
	port          int
	stopTimeoutMS time.Duration
	srvCh         chan error
	options       ServerOptions
	logger        logger.Logger
}

func RunServer(server Server, quit *chan os.Signal) (err error) {
	if server == nil {
		return errors.New("cannot run nil server")
	}

	if server.isLoggerMissing() {
		return errors.New("logger was not found, please provide Logger to Server.AddLogger function")
	}

	signal.Notify(*quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("starting server...")
	serverChannel := server.Start()

	select {
	case osSignal := <-*quit:
		fmt.Printf("received signal [%s], stopping server\n", osSignal.String())
	case serverErr := <-serverChannel:
		fmt.Printf("server start error: %v", serverErr)
	}

	stopServerLogMessage := "server stopped without errors"
	err = server.Stop()
	if err != nil {
		stopServerLogMessage = fmt.Sprintf("server stopped with error: %v", err)
	}

	fmt.Println(stopServerLogMessage)

	return
}
func NewServer(wsConfig *WebServerConfig, router http.Handler, logger logger.Logger, options ServerOptions) Server {
	if logger == nil {
		fmt.Println("logger missing")
	}

	port, stopTimeout := initServerSettings(wsConfig)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	srvCh := make(chan error, 1)

	return &server{
		router:        router,
		httpServer:    httpServer,
		port:          port,
		stopTimeoutMS: stopTimeout,
		srvCh:         srvCh,
		options:       options,
		logger:        logger,
	}
}

func (s *server) Start() (serverChannel chan error) {
	if s.options.UseLogger {
		action, message, customMessage := "ServerStart", "Start", fmt.Sprintf("Starting server on port %d", s.port)
		s.logger.Info(logger.LogMessage{
			Action:        action,
			Message:       message,
			CustomMessage: &customMessage,
		})
	}
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			if s.options.UseLogger {
				action, message, customMessage := "ServerStart", "Start", fmt.Sprintf("Server start error: %v", err)
				s.logger.Error(logger.LogMessage{
					Action:        action,
					Message:       message,
					CustomMessage: &customMessage,
				})
			}
			s.srvCh <- err
		}
	}()

	return s.srvCh
}
func (s *server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.stopTimeoutMS)
	defer cancel()

	//stopHandlers := s.GetStopHandlers()
	//if stopHandlers != nil {
	//	defer func() {
	//		for _, handler := range *stopHandlers {
	//			err := handler.Stop()
	//			if err != nil && s.logger != nil {
	//				message := fmt.Sprintf("Error stopping component: %s", handler.Description)
	//				s.logger.Error(common.LogMessage{
	//					Action:  "StopHandlers",
	//					Message: &message,
	//				})
	//			}
	//		}
	//	}()
	//}

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %v", err)
	}
	return nil
}

func (s *server) AddLogger(l logger.Logger) {
	if l != nil {
		s.logger = l
	}
}

func (s *server) isLoggerMissing() bool {
	return s.options.UseLogger && s.logger == nil
}

func initServerSettings(wsConfig *WebServerConfig) (port int, stopTimeout time.Duration) {
	if wsConfig == nil {
		port = DefaultPort
		stopTimeout = time.Duration(DefaultStopTimeoutMS) * time.Millisecond
	} else {
		port = wsConfig.Port
		stopTimeout = time.Duration(wsConfig.StopTimeoutMS) * time.Millisecond
	}
	return
}
