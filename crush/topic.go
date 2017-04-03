package crush

import (
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
)

// Topic contains all the messages and flight info for a given topic
type Topic struct {
	name        string
	lock        *sync.Mutex
	messages    map[string]*Message
	flight      map[string]*Message
	deadLetterQ chan *Message
}

// CreateTopic will init our defaults for a topic
func CreateTopic(name string, deadLetterQ chan *Message) *Topic {
	t := &Topic{name: name, deadLetterQ: deadLetterQ}
	t.lock = &sync.Mutex{}
	t.messages = make(map[string]*Message)
	t.flight = make(map[string]*Message)
	log.WithFields(log.Fields{"topic": name}).Info("New topic created")
	return t
}

// NewMessage will create a new message on a topic
func (t *Topic) NewMessage(topic, id, value string) *Message {
	msg := NewMessage(t.name, id, value)
	t.lock.Lock()
	t.messages[id] = msg
	t.lock.Unlock()
	log.WithFields(log.Fields{"topic": msg.Topic, "id": msg.ID, "flight": msg.Flight, "attempts": msg.Attempts}).Info("New Message")
	return msg
}

// NewRawMessage creates a new raw message specified by the end user
func (t *Topic) NewRawMessage(msg *Message) *Message {
	t.lock.Lock()
	t.messages[msg.ID] = msg
	t.lock.Unlock()
	log.WithFields(log.Fields{"topic": msg.Topic, "id": msg.ID, "flight": msg.Flight, "attempts": msg.Attempts}).Info("New Message")
	return msg
}

// Messages returns X messages requested from the specific topic
func (t *Topic) Messages(many int) []*Message {
	msgs := make([]*Message, 0)
	for x := 0; x < many; x++ {
		m := t.Message()
		if m != nil {
			msgs = append(msgs, m)
		}
	}
	return msgs
}

// Message returns a single message from the topic
func (t *Topic) Message() *Message {
	t.lock.Lock()
	for id, msg := range t.messages {
		delete(t.messages, id)
		t.flight[id] = msg
		t.lock.Unlock()
		go t.WatchMessage(id, msg)
		return msg
	}
	t.lock.Unlock()
	return nil
}

// WatchMessage follows the topic from lease to completion, seeing if we should add it back to the topic or not
func (t *Topic) WatchMessage(id string, msg *Message) {
	for {
		t.lock.Lock()
		if _, completed := t.flight[id]; !completed {
			t.lock.Unlock()
			break
		}
		flightdur, flighterr := time.ParseDuration(msg.Flight)
		if flighterr != nil {
			flightdur = 5 * time.Minute
		}

		created := time.Unix(msg.Created, 0)
		requeued := time.Unix(msg.Requeued, 0)

		if (msg.Requeued == 0 && time.Now().After(created.Add(flightdur))) || (msg.Requeued != 0 && time.Now().After(requeued.Add(flightdur))) {
			delete(t.flight, id)
			t.lock.Unlock()
			if msg.Attempts > 1 || msg.Attempts == -1 {
				go t.ReQueueMessage(msg)
			} else {
				go t.ExpireMessage(msg)
				if msg.DeadLetter != "" {
					// if deadletter isn't empty, lets send it on it's way ...
					t.deadLetterQ <- msg
				}
			}
			break
		}
		t.lock.Unlock()
		<-time.After(1 * time.Second)
	}
}

// CompleteMessage is an alias to DeleteMessage
func (t *Topic) CompleteMessage(topic, id string) {
	t.DeleteMessage(id)
	log.WithFields(log.Fields{"topic": topic, "id": id}).Info("Message completed")
}

// ExpireMessage is an alias to DeleteMessage
func (t *Topic) ExpireMessage(msg *Message) {
	t.DeleteMessage(msg.ID)
	log.WithFields(log.Fields{"topic": msg.Topic, "id": msg.ID}).Info("Message expired")
}

// ReQueueMessage will put the message back onto the queue so workers can consume it
func (t *Topic) ReQueueMessage(msg *Message) {
	if msg.Attempts != -1 {
		msg.Attempts--
	}
	msg.Requeued = time.Now().Unix()
	t.DeleteMessage(msg.ID)
	t.lock.Lock()
	t.messages[msg.ID] = msg
	t.lock.Unlock()
	log.WithFields(log.Fields{"topic": msg.Topic, "id": msg.ID}).Info("Message requeued")
}

// DeleteMessage removes a message given an id from the topic
func (t *Topic) DeleteMessage(id string) {
	t.lock.Lock()
	delete(t.messages, id)
	delete(t.flight, id)
	t.lock.Unlock()
}
