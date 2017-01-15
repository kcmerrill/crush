package main

import (
	"strings"
	"testing"
)

func TestMessageNewMessage(t *testing.T) {
	m := NewMessage("topic", "id", "value")
	if m.Id != "id" {
		t.Error("id incorrectly set")
	}
	if m.Topic != "topic" {
		t.Error("topic incorrectly set")
	}
	if m.Attempts != 3 {
		t.Error("default attempts should be 3")
	}
}

func TestMessageFormat(t *testing.T) {
	m := NewMessage("topic", "id", "value")
	formatTestOne := m.Format("json")
	if !strings.HasPrefix(formatTestOne, `{"topic":"topic","id":"id","value":"value","created":`) {
		t.Error(formatTestOne + " should be a json formatted string")
	}

	formatTestTwo := m.Format("doesnotexist")
	if formatTestTwo != "" {
		t.Error("Invalid message format should return a blank string")
	}
}
