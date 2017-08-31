package core

import "testing"

func TestStringParam(t *testing.T) {
	s := NewString()

	// set to my email
	s.Set("kcmerrill@gmail.com")

	if email := s.String(); email != "kcmerrill@gmail.com" {
		t.Fatalf("Expected 'kcmerrill@gmail.com', Actual: %s", email)
	}

	if lv := s.List(); len(lv) != 1 {
		t.Fatalf("Expected string ListValue() to return a slice of 1")
	}

	if lv := s.List(); lv[0] != "kcmerrill@gmail.com" {
		t.Fatalf("Expected string ListValue() to return a slice of 1 with [0] = 'kcmerrill@gmail.com'")
	}

	// test reset
	s.Reset()

	if empty := s.String(); empty != "" {
		t.Fatalf("Expected string Reset() to reset the value to an empty string")
	}

	// verify created is not empty
	if s.Created().IsZero() {
		t.Fatalf("Expected NewString() to init String.created to time.Time")
	}

	// verify lastmodified is not empty
	if s.LastModified().IsZero() {
		t.Fatalf("Expected NewString() to init String.modified to time.Time")
	}

	// test many/read writes
	for x := 0; x <= 300; x++ {
		go func() {
			s.Set("asdf")
			s.Reset()
		}()
	}
}
