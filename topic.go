package main

import (
	log "github.com/Sirupsen/logrus"
	"sync"
	"time"
)

type Topic struct {
	name     string
	lock     *sync.Mutex
	messages map[string]*Message
	flight   map[string]*Message
}

func CreateTopic(name string) *Topic {
	t := &Topic{name: name}
	t.lock = &sync.Mutex{}
	t.messages = make(map[string]*Message)
	t.flight = make(map[string]*Message)
	log.WithFields(log.Fields{"topic": name}).Info("New topic created")
	return t
}

func (t *Topic) NewMessage(topic, id, value string) *Message {
	msg := NewMessage(t.name, id, value)
	t.lock.Lock()
	t.messages[id] = msg
	t.lock.Unlock()
	log.WithFields(log.Fields{"topic": topic, "id": id}).Info("New Message")
	return msg
}

func (t *Topic) NewRawMessage(msg *Message) {
	t.lock.Lock()
	t.messages[msg.Id] = msg
	t.lock.Unlock()
	log.WithFields(log.Fields{"topic": msg.Topic, "id": msg.Id}).Info("New Raw Message")
}

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

func (t *Topic) WatchMessage(id string, msg *Message) {
	for {
		<-time.After(1 * time.Second)
		t.lock.Lock()
		if _, completed := t.flight[id]; !completed {
			t.lock.Unlock()
			break
		}
		flightdur, flighterr := time.ParseDuration(msg.Flight)
		if flighterr != nil {
			flightdur = 5 * time.Minute
		}
		if time.Now().After(msg.Requeued.Add(flightdur)) {
			delete(t.flight, id)
			t.lock.Unlock()
			if msg.Attempts > 1 {
				go t.ReQueueMessage(msg)
			} else {
				go t.ExpireMessage(msg)
			}
			break
		}
		t.lock.Unlock()
	}
}

func (t *Topic) CompleteMessage(topic, id string) {
	t.DeleteMessage(id)
	log.WithFields(log.Fields{"topic": topic, "id": id}).Info("Message completed")
}

func (t *Topic) ExpireMessage(msg *Message) {
	t.DeleteMessage(msg.Id)
	log.WithFields(log.Fields{"topic": msg.Topic, "id": msg.Id}).Info("Message expired")
}

func (t *Topic) ReQueueMessage(msg *Message) {
	msg.Attempts--
	msg.Requeued = time.Now()
	t.DeleteMessage(msg.Id)
	t.lock.Lock()
	t.messages[msg.Id] = msg
	t.lock.Unlock()
	log.WithFields(log.Fields{"topic": msg.Topic, "id": msg.Id}).Info("Message requeued")
}

func (t *Topic) DeleteMessage(id string) {
	t.lock.Lock()
	delete(t.messages, id)
	delete(t.flight, id)
	t.lock.Unlock()
}
