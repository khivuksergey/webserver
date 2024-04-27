package webserver

type StopHandler struct {
	Description string
	Stop        func() error
}
