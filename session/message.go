package session

import "strings"

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

func (m Message) GetArgsChannel() string {
	s, _ := m.Args["channel"].(string)

	return strings.TrimSpace(s)
}

func (m Message) GetArgsContent() any {
	return m.Args["content"]
}
