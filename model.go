package webserver

type ServerOptions struct{}

type StopHandler struct {
	Description string
	Stop        func() error
}
