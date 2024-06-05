package session

import "fmt"

type MessageCommandType string

const (
	MessageCommandTypeJoin      = MessageCommandType("join")
	MessageCommandTypeLeave     = MessageCommandType("leave")
	MessageCommandTypeBroadcast = MessageCommandType("broadcast")
)

type Message struct {
	Command MessageCommandType `json:"command"`
	Args    map[string]any     `json:"args"`
}

func (m Message) GetArgsChannels() (result []string) {
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

func (m Message) GetArgsContent() any {
	return m.Args["content"]
}
