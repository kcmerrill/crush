package core

import (
	"sync"
	"time"
)

// NewBool inits the Date struct
func NewBool() *Bool {
	return &Bool{
		CreatedDate: time.Now(),
		Modified:    time.Now(),
		lock:        &sync.Mutex{},
	}
}

// Bool property type
type Bool struct {
	CreatedDate time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
	Value       bool      `json:"value"`
	lock        *sync.Mutex
}

// Reset Bool to now
func (b *Bool) Reset() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.Value = false
}

// Set the value to be something
func (b *Bool) Set(something interface{}) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.Value = something.(bool)
	b.Modified = time.Now()
}

// LastModified returs the last modified time
func (b *Bool) LastModified() time.Time {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.Modified
}

// Created returns the created time
func (b *Bool) Created() time.Time {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.CreatedDate
}

// String returns the Bools value
func (b *Bool) String() string {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.Value {
		return "true"
	}
	return "false"
}

// Int returns the Bools value
func (b *Bool) Int() int {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.Value {
		return 1
	}
	return 0
}

// List converts Bool to a list. Why? don't ask me ...
func (b *Bool) List() []string {
	b.lock.Lock()
	defer b.lock.Unlock()
	return []string{b.String()}
}

// Add not implemented
func (b *Bool) Add(something interface{}) {}

// Remove not implemented
func (b *Bool) Remove(something interface{}) {}

// Bool returns the value of bool
func (b *Bool) Bool() bool {
	return b.Value
}

// Type returns the type of this property
func (b *Bool) Type() string {
	return "bool"
}
