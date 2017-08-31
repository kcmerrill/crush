package core

import "testing"
import "time"

func TestSherlockEntity(t *testing.T) {
	s := New(100)
	s.Entity("kcmerrill@gmail.com").NewProperty("username", "string").Set("themayor")

	if name := s.Entity("kcmerrill@gmail.com").Property("username").String(); name != "themayor" {
		t.Fatalf("Expected 'themayor', Actual: '%s'", name)
	}

	// make sure the entity creation time isn't zero
	if s.Entity("kcmerrill@gmail.com").Created().IsZero() {
		t.Fatalf("Created sould not be a zero time.Time")
	}

	// lets play with the counter now
	e := s.Entity("kcmerrill@gmail.com")
	e.NewProperty("counter", "int").Set(1000)

	if i := s.Entity("kcmerrill@gmail.com").Property("counter").Int(); i != 1000 {
		t.Fatalf("Was expecting 'counter' to be 1000")
	}

	// Add to it
	e.Property("counter").Add(100)

	if i := s.Entity("kcmerrill@gmail.com").Property("counter").Int(); i != 1100 {
		t.Fatalf("Was expecting 'counter' to be 1100")
	}

	// quick, add an event
	e.Event("clicked on button 'A'")
	e.Event("clicked on button 'B'")

	if e.Events[0].Description != "clicked on button 'A'" {
		t.Fatalf("Was Expecting button a to be clicked!")
	}
}

func TestShortValueCreators(t *testing.T) {
	s := New(100)
	s.E("kcmerrill").I("counter").Set(10)
	if count := s.E("kcmerrill").I("counter").Int(); count != 10 {
		t.Fatalf("Expected 10, Actual: %d", count)
	}

	_created := s.E("kcmerrill").D("_created").Int()
	now := time.Now().Unix()
	if _created != int(now) {
		t.Fatalf("Expected _created to be now()")
	}

	s.E("bingowas").S("his").Set("nameo")
	nameo := s.E("bingowas").S("his").String()
	if nameo != "nameo" {
		t.Fatalf("Expected 'nameo', Actual '%s'", nameo)
	}

	// while we are here, lets test resets
	s.E("kcmerrill").I("counter").Reset()
	val := s.E("kcmerrill").I("counter").Int()
	if val != 0 {
		t.Fatalf("Expected after reset() counter to be 0")
	}

	s.E("bingowas").S("his").Reset()
	empty := s.E("bingowas").S("his").String()
	if empty != "" {
		t.Fatalf("Expected '', Actual '%s'", empty)
	}
}

func TestEntityHas(t *testing.T) {
	s := New(100)
	s.E("kcmerrill@gmail.com")
	if s.E("kcmerrill@gmail.com").Has("username") {
		t.Fatalf("We didn't set an email on this entity ... should not exist")
	}

	s.E("kcmerrill@gmail.com").S("username").Set("kcmerrill")

	if !s.E("kcmerrill@gmail.com").Has("username") {
		t.Fatalf("username should now be set on this entity")
	}
}
