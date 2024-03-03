package logger

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Logger interface {
	Debug(LogMessage)
	Info(LogMessage)
	Warn(LogMessage)
	Error(LogMessage)
	Fatal(LogMessage)
}

type LogMessage struct {
	time             time.Time
	Action           string
	Message          string
	CustomMessage    *string
	UserId           *int
	RequestGuid      *uuid.UUID
	Data             any
	AdditionalFields *map[string]any
}

func (m *LogMessage) String() (formatted string) {
	m.time = time.Now()
	formatted += m.time.Format("02/01/2006 15:04:05.000000")

	// Log level
	formatted += " [%s]"

	if m.Action != "" {
		formatted += fmt.Sprintf(" Action: %s,", m.Action)
	}

	if m.Message != "" {
		formatted += fmt.Sprintf(" Message: %s,", m.Message)
	}

	if m.CustomMessage != nil {
		formatted += fmt.Sprintf(" CustomMessage: %s,", *m.CustomMessage)
	}

	if m.UserId != nil {
		formatted += fmt.Sprintf(" UserId: %d,", *m.UserId)
	}

	if m.RequestGuid != nil {
		formatted += fmt.Sprintf(" RequestGuid: %s,", m.RequestGuid.String())
	}

	if m.Data != nil {
		formatted += fmt.Sprintf(" Data: %v,", m.Data)
	}

	if m.AdditionalFields != nil {
		formatted += fmt.Sprintf(" AdditionalFields: %v", *m.AdditionalFields)
	}

	if len(formatted) > 0 && formatted[len(formatted)-1] == ',' {
		formatted = formatted[:len(formatted)-1]
	}

	formatted += "\n"

	return
}
