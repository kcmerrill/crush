package core

import (
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// Topic contains all the messages and flight info for a given topic
type Topic struct {
	name     string
	lock     *sync.Mutex
	messages map[string]*Message
	flight   map[string]*Message
	q        *Q
}

// CreateTopic will init our defaults for a topic
func CreateTopic(name string, q *Q) *Topic {
	t := &Topic{name: name, q: q}
	t.lock = &sync.Mutex{}
	t.messages = make(map[string]*Message)
	t.flight = make(map[string]*Message)
	log.WithFields(log.Fields{"topic": name}).Info("New topic created")
	q.Stat.E("_system").I("topics").Add(1)
	return t
}

// NewMessage will create a new message on a topic
func (t *Topic) NewMessage(topic, id, value string) *Message {
	msg := NewMessage(t.name, id, value)
	t.lock.Lock()
	t.messages[id] = msg
	t.lock.Unlock()
	t.q.Stat.E("_system").I("new-messages").Add(1)
	t.q.Stat.E(t.name).I("new-messages").Add(1)
	log.WithFields(log.Fields{"topic": msg.Topic, "id": msg.ID, "flight": msg.Flight, "attempts": msg.Attempts}).Info("New Message")
	return msg
}

// NewRawMessage creates a new raw message specified by the end user
func (t *Topic) NewRawMessage(msg *Message) *Message {
	t.lock.Lock()
	t.messages[msg.ID] = msg
	t.lock.Unlock()
	t.q.Stat.E("_system").I("new-messages").Add(1)
	t.q.Stat.E(t.name).I("new-messages").Add(1)
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
		t.q.Stat.E("_system").I("processed-messages").Add(1)
		t.q.Stat.E(t.name).I("processed-messages").Add(1)
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
				// we already checked, but lets check again. Safety pumpking safety ...
				if msg.DeadLetter != "" {
					for _, newTopic := range strings.Split(msg.DeadLetter, " ") {
						// dont' really want to increment this one...
						t.q.Stat.E("_system").I("new-messages").Add(-1)
						t.q.Stat.E(t.name).I("new-messages").Add(-1)
						log.WithFields(log.Fields{"previous-topic": t.name, "id": msg.ID, "new-topic": newTopic}).Info("Moving message to '" + newTopic + "'")
					}
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
	t.q.Stat.E("_system").I("completed-messages").Add(1)
	t.q.Stat.E(t.name).I("completed-messages").Add(1)
}

// ExpireMessage is an alias to DeleteMessage
func (t *Topic) ExpireMessage(msg *Message) {
	t.DeleteMessage(msg.ID)
	log.WithFields(log.Fields{"topic": msg.Topic, "id": msg.ID}).Info("Message expired")
	t.q.Stat.E("_system").I("expired-messages").Add(1)
	t.q.Stat.E(t.name).I("expired-messages").Add(1)
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
	t.q.Stat.E("_system").I("requeued-messages").Add(1)
	t.q.Stat.E(t.name).I("requeued-messages").Add(1)
}

// DeleteMessage removes a message given an id from the topic
func (t *Topic) DeleteMessage(id string) {
	t.lock.Lock()
	delete(t.messages, id)
	delete(t.flight, id)
	t.lock.Unlock()
	t.q.Stat.E("_system").I("deleted-messages").Add(1)
	t.q.Stat.E(t.name).I("deleted-messages").Add(1)
}
