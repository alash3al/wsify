package main

// the message that will be polled from redis
type Message struct {
	Payload interface{} `json:"payload,omitempty"`
	To      []string    `json:"to,omitempty"`
	Topic   string      `json:"channel,omitempty"`
	Time    int64       `json:"time,omitempty"`
}

// checks whether the specified user is allowed
// to recieve this message or not
func (m *Message) IsUserAllowed(u string) bool {
	if len(m.To) < 1 {
		return true
	}
	for _, v := range m.To {
		if v == u {
			return true
		}
	}
	return false
}
