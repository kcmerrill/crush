package main

import (
	"strings"
	"testing"
)

func TestMessageNewMessage(t *testing.T) {
	m := NewMessage("topic", "id", "value")
	if m.ID != "id" {
		t.Error("id incorrectly set")
	}
	if m.Topic != "topic" {
		t.Error("topic incorrectly set")
	}
	if m.Attempts != 3 {
		t.Error("default attempts should be 3")
	}
}

func TestMessageString(t *testing.T) {
	m := NewMessage("topic", "id", "value")
	formatTestOne := m.String()
	if !strings.HasPrefix(formatTestOne, `{"topic":"topic","id":"id","value":"value","created":`) {
		t.Error(formatTestOne + " should be a json formatted string")
	}
}
