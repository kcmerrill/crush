package main

import (
	"testing"
)

func TestCreateQ(t *testing.T) {
	q := CreateQ()
	if len(q.topics) != 0 {
		t.Error("topics should be 0 in length")
	}
}

func TestQNewMessage(t *testing.T) {
	q := CreateQ()
	// first test create topic
	q.NewMessage("test", "id", "value")
	if _, exists := q.topics["test"]; !exists {
		t.Error("Topic test was not created")
	}
	// next, test create a new message on an existing topic
	q.NewMessage("test", "id2", "value2")
	if len(q.topics["test"].messages) != 2 {
		t.Error("Two messages were added. Two should be in messages[]")
	}
}

func TestQMessage(t *testing.T) {
	q := CreateQ()
	// Lets fetch the message
	q.NewMessage("test", "id", "kcwashere")
	m := q.Message("test")
	if m.Value != "kcwashere" {
		t.Error("Unable to fetch a valid message")
	}
	// now, fetch a message(should be blank)
	empty := q.Message("test")
	if empty != nil {
		t.Error("Message visible when it shouldn't have been")
	}
}

func TestQMessages(t *testing.T) {
	q := CreateQ()
	// Topic does not exist, lets test
	notopic := q.Messages("notopic", 10)
	if len(q.topics["notopic"].messages) != 0 {
		t.Error("The topic should exist, and there should be no messages in it")
	}

	if len(notopic) != 0 {
		t.Error("There should not be any messages in a non existant topic")
	}

	// Lets create 2 messages
	q.NewMessage("test", "id", "kcwashere")
	q.NewMessage("test", "id2", "kcwashere2")
	if len(q.topics["test"].messages) != 2 {
		t.Error("Should be 2 messages available to consume")
	}
	q.NewMessage("test", "id", "kcwashere")
	if len(q.topics["test"].messages) != 2 {
		t.Error("Should be 2 messages available to consume, even with same id")
	}
	// Array should be 2
	msgs := q.Messages("test", 10)
	if len(msgs) != 2 {
		t.Error("q.Messages() should return only two messages")
	}
}

func TestQComplete(t *testing.T) {
	// lets test topic that doesn't exist
	q := CreateQ()
	if _, exists := q.topics["test"]; exists {
		t.Error("Test queue should not exist yet")
	}
	// Complete a message on a queue that doesn't exist
	tm := &Message{
		Id: "id",
	}
	q.Complete("test", tm.Id)

	// lets try adding the message, and trying again
	q.NewMessage("test", "id", "value")
	if len(q.topics["test"].messages) != 1 {
		t.Error("test topic should have 1 message")
	}

	// grab the message
	m := q.Message("test")
	q.Complete("test", m.Id)
	if len(q.topics["test"].messages) != 0 {
		t.Error("test topic should have 0 message after completion")
	}
}

func TestDeleteMessage(t *testing.T) {
	q := CreateQ()
	q.NewMessage("test", "id", "value")
	if len(q.topics["test"].messages) != 1 {
		t.Error("test topic should have 1 message")
	}
	q.Delete("test", "id")
	if len(q.topics["test"].messages) != 0 {
		t.Error("Deleting of messgaes is not working")
	}
}
