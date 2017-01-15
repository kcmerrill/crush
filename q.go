package main

import (
	"sync"
)

type Q struct {
	lock   *sync.Mutex
	topics map[string]*Topic
}

func CreateQ() *Q {
	q := &Q{}
	q.topics = make(map[string]*Topic)
	q.lock = &sync.Mutex{}
	return q
}

func (q *Q) NewMessage(topic, key, value string) *Message {
	q.lock.Lock()
	var m *Message
	if _, exists := q.topics[topic]; exists {
		m = q.topics[topic].NewMessage(topic, key, value)
	} else {
		q.topics[topic] = CreateTopic(topic)
		m = q.topics[topic].NewMessage(topic, key, value)
	}
	q.lock.Unlock()
	return m
}

func (q *Q) Message(topic string) *Message {
	q.lock.Lock()
	var m *Message
	if _, exists := q.topics[topic]; exists {
		m = q.topics[topic].Message()
	} else {
		q.topics[topic] = CreateTopic(topic)
		m = q.topics[topic].Message()
	}
	q.lock.Unlock()
	return m
}

func (q *Q) Messages(topic string, count int) []*Message {
	q.lock.Lock()
	var m []*Message
	if _, exists := q.topics[topic]; exists {
		m = q.topics[topic].Messages(count)
	} else {
		q.topics[topic] = CreateTopic(topic)
		m = q.topics[topic].Messages(count)
	}
	q.lock.Unlock()
	return m
}

func (q *Q) Complete(topic, id string) {
	q.lock.Lock()
	if _, exists := q.topics[topic]; exists {
		q.topics[topic].CompleteMessage(topic, id)
	} else {
		q.topics[topic] = CreateTopic(topic)
		q.topics[topic].CompleteMessage(topic, id)
	}
	q.lock.Unlock()
}

func (q *Q) Delete(topic, id string) {
	q.lock.Lock()
	if _, exists := q.topics[topic]; exists {
		q.topics[topic].DeleteMessage(id)
	} else {
		q.topics[topic] = CreateTopic(topic)
		q.topics[topic].DeleteMessage(id)
	}
	q.lock.Unlock()
}
