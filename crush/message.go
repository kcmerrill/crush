package crush

import (
	"encoding/json"
	"time"
)

// NewMessage returns a default init'd message
func NewMessage(topic, id, value string) *Message {
	return &Message{
		ID:         id,
		Topic:      topic,
		Value:      value,
		Attempts:   3,
		Created:    time.Now().Unix(),
		Flight:     "5m",
		DeadLetter: "",
	}
}

// Message contains information needed to create a valid q message
type Message struct {
	Topic      string `json:"topic"`
	ID         string `json:"id"`
	Value      string `json:"value"`
	Created    int64  `json:"created"`
	Requeued   int64  `json:"requeued"`
	Flight     string `json:"flight"`
	Attempts   int    `json:"attempts"`
	DeadLetter string `json:"dead_letter"`
}

func (m *Message) String() string {
	str, err := json.Marshal(m)
	if err != nil {
		return ""
	}

	return string(str)
}
