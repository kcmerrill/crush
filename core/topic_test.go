package core

import (
	"testing"
	"time"
)

func TestCreateTopic(t *testing.T) {
	to := CreateTopic("bingowashisnameo", make(chan *Message))
	if to.name != "bingowashisnameo" {
		t.Error("Name not set properly")
	}

	if len(to.messages) != 0 {
		t.Error("messages not initialized to 0")
	}

	if len(to.flight) != 0 {
		t.Error("flight not initialized to 0")
	}
}

func TestTopicNewMessage(t *testing.T) {
	to := CreateTopic("test", make(chan *Message))
	// message count should be zero
	if len(to.messages) != 0 {
		t.Error("messages should be zero")
	}
	// add a message
	to.NewMessage("test", "id", "value")
	// count the new message
	if len(to.messages) != 1 {
		t.Error("messages should be 1")
	}
	// grab the message and verify the contents
	m := to.Message()
	if m.Value != "value" {
		t.Error("Unable to get the proper value")
	}
}

func TestTopicNewRawMessage(t *testing.T) {
	// same as new message, except raw only
	topic := CreateTopic("bingowashisnameo", make(chan *Message))
	msg := &Message{
		ID:    "id",
		Value: "woot",
		Topic: "bingo",
	}
	topic.NewRawMessage(msg)
	m := topic.Message()
	if m.Value != "woot" {
		t.Error("Unable to create new raw message")
	}
}

func TestTopicMessages(t *testing.T) {
	// Grab a bunch of messages
	topic := CreateTopic("bingowashisnameo", make(chan *Message))
	topic.NewMessage("bingowashisnameo", "id", "value")
	topic.NewMessage("bingowashisnameo", "id2", "value3")
	topic.NewMessage("bingowashisnameo", "id3", "value3")

	msgs := topic.Messages(100)

	// even though we asked for 100, there are only 3 current messages
	if len(msgs) != 3 {
		t.Error("messages() should only return 3 messages")
	}

	// now, really quick, lets grab another hundred. Should be empty
	msgs = topic.Messages(100)
	if len(msgs) != 0 {
		t.Error("There should not be any messages valid at this point")
	}
}

func TestWatchMessage(t *testing.T) {
	dl := make(chan *Message)
	topic := CreateTopic("bleh", dl)
	/* Create a message */
	m := NewMessage("bleh", "id", "testwatchmessage()")
	m.DeadLetter = "dead-letter"
	m.Attempts = 2
	m.Flight = "100ms"
	topic.NewRawMessage(m)

	/* Grab the message, verify it's contents */
	msgAttemptOne := topic.Message()
	if msgAttemptOne.Value != "testwatchmessage()" {
		t.Error(msgAttemptOne.Value + " is not the value for the message we were expecting")
	}

	if len(topic.flight) != 1 {
		t.Error("The message should be in flight but it's not in flight[]")
	}

	<-time.After(1*time.Second + 10*time.Millisecond)

	// Grab the message ... should be "requeued" aka attempts should be limited now
	msgAttemptTwo := topic.Message()
	if msgAttemptTwo.Value != "testwatchmessage()" {
		t.Error(msgAttemptTwo.Value + " is not the value for the message we were expecting")
	}

	<-time.After(1*time.Second + 10*time.Millisecond)

	msgAttemptThree := topic.Message()

	if msgAttemptThree != nil {
		t.Error("The third attempt should be null")
	}

	// because we setup a deadletter queue, lets process it and see what we got!
	deaddeaddead := <-dl
	if deaddeaddead.DeadLetter != "dead-letter" {
		t.Fatalf("Expected: 'dead-letter', actual: '%s'", deaddeaddead.DeadLetter)
	}

	if deaddeaddead.ID != "id" {
		t.Fatalf("Expected: 'id', actual: '%s'", deaddeaddead.ID)
	}

	if deaddeaddead.Value != "testwatchmessage()" {
		t.Fatalf("Expected: 'testwatchmessage()', actual: '%s'", deaddeaddead.Value)
	}
}
