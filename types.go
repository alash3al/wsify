package main

import "fmt"

type WebSocketMessageCommandType string

const (
	WebSocketMessageCommandTypeJoin      = WebSocketMessageCommandType("join")
	WebSocketMessageCommandTypeLeave     = WebSocketMessageCommandType("leave")
	WebSocketMessageCommandTypeBroadcast = WebSocketMessageCommandType("broadcast")
)

type WebSocketMessage struct {
	Command WebSocketMessageCommandType `json:"command"`
	Args    map[string]any              `json:"args"`
}

func (m WebSocketMessage) GetAvailableChannels() (result []string) {
	if m.Args["channels"] == nil {
		return []string{}
	}

	switch val := m.Args["channels"].(type) {
	case []string:
		result = val
	case []any:
		for _, v := range val {
			result = append(result, fmt.Sprintf("%v", v))
		}
	}

	return
}

func (m WebSocketMessage) GetAvailableWebSocketMessage() any {
	return m.Args["message"]
}
