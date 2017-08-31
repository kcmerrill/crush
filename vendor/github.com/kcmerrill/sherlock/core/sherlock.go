package core

import (
	"encoding/json"
	"sync"
	"time"
)

// Entity holds our entity information
type Entity struct {
	ID         string `json:"id"`
	lock       *sync.Mutex
	Properties map[string]Property `json:"properties"`
	Events     []*Event            `json:"events"`
	maxEvents  int
}

// Event object tracks a specific event
type Event struct {
	Created     time.Time `json:"created"`
	Description string    `json:"description"`
}

// NewEvent creates a new event
func NewEvent(msg string) *Event {
	return &Event{
		Created:     time.Now(),
		Description: msg,
	}
}

// Property will return an entities property
func (e *Entity) Property(name string) Property {
	e.lock.Lock()
	defer e.lock.Unlock()

	// no error checking? YOLO
	return e.Properties[name]
}

// Event will create a new event for an entitiy
func (e *Entity) Event(msg string) {
	e.lock.Lock()
	defer e.lock.Unlock()

	event := NewEvent(msg)

	e.Events = append(e.Events, event)

	// TODO: Make this value customizable
	if len(e.Events) > 50 {
		e.Events = e.Events[1:len(e.Events)]
	}
}

// I will create a new int property if it doesn't exist
func (e *Entity) I(name string) Property {
	e.lock.Lock()
	p, exists := e.Properties[name]
	// unlock everything
	e.lock.Unlock()
	if exists {
		return p
	}
	return e.NewProperty(name, "int")
}

// B will create a new bool property if it doesn't exist
func (e *Entity) B(name string) Property {
	e.lock.Lock()
	p, exists := e.Properties[name]
	// unlock everything
	e.lock.Unlock()
	if exists {
		return p
	}
	return e.NewProperty(name, "bool")
}

// S will create a new string property if it doesn't exist
func (e *Entity) S(name string) Property {
	e.lock.Lock()
	p, exists := e.Properties[name]
	// unlock everything
	e.lock.Unlock()
	if exists {
		return p
	}
	return e.NewProperty(name, "string")
}

// D will create a new string property if it doesn't exist
func (e *Entity) D(name string) Property {
	e.lock.Lock()
	p, exists := e.Properties[name]
	// unlock everything
	e.lock.Unlock()
	if exists {
		return p
	}
	return e.NewProperty(name, "date")
}

// NewProperty will create and return a new property
func (e *Entity) NewProperty(name, param string) Property {
	e.lock.Lock()
	defer e.lock.Unlock()

	// see if it already exists
	if _, exists := e.Properties[name]; exists {
		return e.Properties[name]
	}

	var p Property
	switch param {
	case "int":
		p = NewInt()
		break
	case "date":
		p = NewDate()
		break
	case "bool":
		p = NewBool()
	case "string":
		fallthrough
	default:
		p = NewString()
	}

	e.Properties[name] = p
	return e.Properties[name]
}

// Has deterines if an entity has a property or not
func (e *Entity) Has(property string) bool {
	e.lock.Lock()
	defer e.lock.Unlock()
	_, exists := e.Properties[property]
	return exists
}

// Created returns the entity creation date(aka the _created param)
func (e *Entity) Created() time.Time {
	created := e.Property("_created").Int()
	return time.Unix(int64(created), 0)
}

// String returns the entity as a string
func (e *Entity) String() (string, error) {
	j, err := json.Marshal(e)
	if err == nil {
		return string(j), err
	}
	return "", err
}

// Property can be multiple things ...
type Property interface {
	Reset()
	Add(something interface{})
	Remove(something interface{})
	Set(something interface{})
	String() string
	Int() int
	List() []string
	Bool() bool
	LastModified() time.Time
	Created() time.Time
	Type() string
}

// Sherlock struct
type Sherlock struct {
	AuthToken string
	lock      *sync.Mutex
	Entities  map[string]*Entity `json:"entities"`
	maxEvents int
}

// E is shorthand for Entity
func (s *Sherlock) E(id string) *Entity {
	return s.Entity(id)
}

// Entity returns a string entity, if none exist, creates one
func (s *Sherlock) Entity(id string) *Entity {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, exists := s.Entities[id]; !exists {
		// we need to create a blank entity
		s.Entities[id] = NewEntity(id, s.maxEvents)
	}

	return s.Entities[id]
}

// New returns a newly initialized sherlock
func New(maxEvents int) *Sherlock {
	s := &Sherlock{maxEvents: maxEvents}
	s.lock = &sync.Mutex{}
	s.Entities = make(map[string]*Entity)
	return s
}

// NewEntity returns a new entity
func NewEntity(id string, maxEvents int) *Entity {
	e := &Entity{ID: id, maxEvents: maxEvents}
	e.Properties = make(map[string]Property)
	e.lock = &sync.Mutex{}
	e.Events = make([]*Event, 0)
	e.NewProperty("_created", "date").Set(time.Now())
	return e
}
