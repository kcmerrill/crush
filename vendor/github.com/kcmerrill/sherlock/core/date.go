package core

import (
	"sync"
	"time"
)

// NewDate inits the Date struct
func NewDate() *Date {
	return &Date{
		CreatedDate: time.Now(),
		Modified:    time.Now(),
		lock:        &sync.Mutex{},
	}
}

// Date property type
type Date struct {
	CreatedDate time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
	Value       time.Time `json:"value"`
	lock        *sync.Mutex
}

// Reset Date to now
func (d *Date) Reset() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.Value = time.Now()
}

// Set the value to be something
func (d *Date) Set(something interface{}) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.Value = something.(time.Time)
	d.Modified = time.Now()
}

// LastModified returs the last modified time
func (d *Date) LastModified() time.Time {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.Modified
}

// Created returns the created time
func (d *Date) Created() time.Time {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.CreatedDate
}

// String returns the Dates value
func (d *Date) String() string {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.Value.String()
}

// Int returns the Dates value
func (d *Date) Int() int {
	d.lock.Lock()
	defer d.lock.Unlock()
	return int(d.Value.Unix())
}

// List converts Date to a list
func (d *Date) List() []string {
	d.lock.Lock()
	defer d.lock.Unlock()
	return []string{d.Value.String()}
}

// Add increments value by something
func (d *Date) Add(something interface{}) {
	d.lock.Lock()
	defer d.lock.Unlock()
	by := something.(string)
	if dur, err := time.ParseDuration(by); err == nil {
		d.Value.Add(dur)
		d.Modified = time.Now()
	}
}

// Remove not implemented
func (d *Date) Remove(something interface{}) {}

// Bool returns if zero()
func (d *Date) Bool() bool {
	return d.Value.IsZero()
}

// Type returns the type of this property
func (d *Date) Type() string {
	return "date"
}
