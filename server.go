package webserver

import (
	"errors"
	"fmt"
	"github.com/khivuksergey/webserver/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	Start() chan error
	Stop() error
	AddLogger(logger.Logger)
	AddStopHandlers(*[]StopHandler)
	IsLoggerMissing() bool
}

func RunServer(server Server, quit chan os.Signal) (err error) {
	if server == nil {
		return errors.New("cannot run nil server")
	}

	if server.IsLoggerMissing() {
		return errors.New("logger was not found, please provide Logger to Server.AddLogger function")
	}

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Printf("%s [SERVER] STARTING SERVER\n", now())
	serverChannel := server.Start()

	select {
	case osSignal := <-quit:
		fmt.Printf("\n%s [SERVER] RECEIVED SIGNAL [%s], STOPPING SERVER\n", now(), osSignal.String())
	case serverErr := <-serverChannel:
		fmt.Printf("%s [SERVER] SERVER START ERROR: %v", now(), serverErr)
	}

	stopServerLogMessage := "SERVER STOPPED WITHOUT ERRORS"
	err = server.Stop()
	if err != nil {
		stopServerLogMessage = fmt.Sprintf("SERVER STOPPED WITH ERROR: %v", err)
	}

	fmt.Printf("%s [SERVER] %s\n", now(), stopServerLogMessage)

	return
}

func now() string {
	return time.Now().Format("02/01/2006 15:04:05.000000")
}
