package logger

import (
	"fmt"
	"strings"
	"time"
)

type LogMessage struct {
	time             time.Time
	Action           string
	Message          string
	CustomMessage    *string
	UserId           *uint64
	Data             any
	AdditionalFields *map[string]any
	RequestUuid      string
}

func (m *LogMessage) String() string {
	var builder strings.Builder

	m.time = time.Now()

	builder.WriteString(m.time.Format("02/01/2006 15:04:05.000000"))

	// Log level
	builder.WriteString(" [%s]")

	if m.Action != "" {
		builder.WriteString(fmt.Sprintf(" Action: %s,", m.Action))
	}

	if m.Message != "" {
		builder.WriteString(fmt.Sprintf(" Message: %s,", m.Message))
	}

	if m.CustomMessage != nil {
		builder.WriteString(fmt.Sprintf(" CustomMessage: %s,", *m.CustomMessage))
	}

	if m.UserId != nil {
		builder.WriteString(fmt.Sprintf(" UserId: %d,", *m.UserId))
	}

	if m.Data != nil {
		builder.WriteString(fmt.Sprintf(" Data: %v,", m.Data))
	}

	if m.AdditionalFields != nil {
		builder.WriteString(fmt.Sprintf(" AdditionalFields: %v,", *m.AdditionalFields))
	}

	builder.WriteString(fmt.Sprintf(" RequestUuid: %s\n", m.RequestUuid))

	return builder.String()
}
