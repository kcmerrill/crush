package main

import (
	"encoding/json"
	"time"
)

func NewMessage(topic, id, value string) *Message {
	return &Message{
		Id:       id,
		Topic:    topic,
		Value:    value,
		Attempts: 3,
		Created:  time.Now().Unix(),
		Requeued: time.Now().Unix(),
		Flight:   "5m",
	}
}

type Message struct {
	Topic    string `json:"topic"`
	Id       string `json:"id"`
	Value    string `json:"value"`
	Created  int64    `json:"created"`
	Requeued int64    `json:"requeued"`
	Flight   string `json:"flight"`
	Attempts int    `json:"attempts"`
}

func (m *Message) Format(which string) string {
	if which == "json" {
		str, err := json.Marshal(m)
		if err != nil {
			return ""
		} else {
			return string(str)
		}
	}
	return ""
}
