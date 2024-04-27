package webserver

type StopHandler struct {
	Description string
	Stop        func() error
}

func NewStopHandler(description string, stop func() error) *StopHandler {
	return &StopHandler{
		Description: description,
		Stop:        stop,
	}
}
