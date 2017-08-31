package core

import (
	"sync"
	"time"
)

// NewString inits the string struct
func NewString() *String {
	return &String{
		CreatedDate: time.Now(),
		Modified:    time.Now(),
		lock:        &sync.Mutex{},
	}
}

// String property type
type String struct {
	CreatedDate time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
	Value       string    `json:"value"`
	lock        *sync.Mutex
}

// Reset string to ""
func (s *String) Reset() {
	s.lock.Lock()
	s.Value = ""
	s.lock.Unlock()
}

// Set the value to be something
func (s *String) Set(something interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Value = something.(string)
	s.Modified = time.Now()
}

// LastModified returs the last modified time
func (s *String) LastModified() time.Time {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Modified
}

// Created returns the created time
func (s *String) Created() time.Time {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.CreatedDate
}

// String returns the strings value
func (s *String) String() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Value
}

// Int not used
func (s *String) Int() int {
	return 0
}

// List converts string to a list
func (s *String) List() []string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return []string{s.Value}
}

// Add not implemented
func (s *String) Add(something interface{}) {}

// Remove not implemented
func (s *String) Remove(something interface{}) {}

// Bool returns a bool if the string is set or not
func (s *String) Bool() bool {
	if s.Value == "" {
		return false
	}
	return true
}

// Type returns the type of this property
func (s *String) Type() string {
	return "string"
}
