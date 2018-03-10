package pubsub

type Message struct {
	topic     string
	payload   interface{}
	createdAt int64
}

// to return the topic of the current message
func (m *Message) GetTopic() string {
	return m.topic
}

// to get the payload of the current message
func (m *Message) GetPayload() interface{} {
	return m.payload
}

// get the creation time of this message
func (m *Message) GetCreatedAt() int64 {
	return m.createdAt
}
