package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/karkulevskiy/shareen/src/internal/ws"
	"github.com/karkulevskiy/shareen/src/internal/ws/events"
)

type ClientHelp struct {
	ws.Client
}

// TestAskForVideoTiming tests AskForVideoTiming
func TestAskForVideoTiming(t *testing.T) {
	tests := []struct {
		Name     string
		Request  string
		Response ws.Event
	}{
		{
			Name:     "Valid test",
			Request:  "Max",
			Response: helpCreateEvent(events.AskVideoTimingEvent{Login: "Max"}, ws.EventGetVideoTiming),
		},
		{
			Name:     "Invalid test",
			Request:  "Bob",
			Response: helpCreateEvent(events.AskVideoTimingEvent{Login: "Bob"}, ws.EventDisconnect),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ws.AskForVideoTiming(tt.Request, &ws.Client{})
		})
	}
}

func helpCreateEvent(event any, eventType string) ws.Event {
	payload, _ := json.Marshal(event)
	return ws.CreateEvent(http.StatusOK, eventType, payload)
}
