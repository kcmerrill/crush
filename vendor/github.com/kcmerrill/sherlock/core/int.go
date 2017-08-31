package core

import (
	"strconv"
	"sync"
	"time"
)

// NewInt inits the Int struct
func NewInt() *Int {
	return &Int{
		CreatedDate: time.Now(),
		Modified:    time.Now(),
		lock:        &sync.Mutex{},
	}
}

// Int property type
type Int struct {
	CreatedDate time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
	Value       int       `json:"value"`
	lock        *sync.Mutex
}

// Reset Int to ""
func (i *Int) Reset() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.Value = 0
}

// Set the value to be something
func (i *Int) Set(something interface{}) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.Value = something.(int)
	i.Modified = time.Now()
}

// LastModified returs the last modified time
func (i *Int) LastModified() time.Time {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.Modified
}

// Created returns the created time
func (i *Int) Created() time.Time {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.CreatedDate
}

// String returns the Ints value
func (i *Int) String() string {
	i.lock.Lock()
	defer i.lock.Unlock()
	return strconv.Itoa(i.Value)
}

// Int returns the ints value
func (i *Int) Int() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.Value
}

// List converts Int to a list
func (i *Int) List() []string {
	i.lock.Lock()
	defer i.lock.Unlock()
	return []string{strconv.Itoa(i.Value)}
}

// Add increments value by something
func (i *Int) Add(something interface{}) {
	i.lock.Lock()
	defer i.lock.Unlock()
	by := something.(int)
	i.Value += by
	i.Modified = time.Now()
}

// Bool returns true/false
func (i *Int) Bool() bool {
	if i.Value > 0 {
		return true
	}
	return false
}

// Remove removes X from the int
func (i *Int) Remove(something interface{}) {
	i.lock.Lock()
	defer i.lock.Unlock()
	by := something.(int)
	i.Value -= by
	i.Modified = time.Now()
}

// Type returns the type of this property
func (i *Int) Type() string {
	return "int"
}
