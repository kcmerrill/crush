package core

import "testing"

func TestBool(t *testing.T) {
	s := New(100)
	b := s.E("kcmerrill").B("awesome")

	// default, should be false :(
	if b.Bool() {
		t.Fatalf("By default, bool should be false")
	}

	b.Set(true)
	if !b.Bool() {
		t.Fatalf("bool was set to false, should be false")
	}

	if b.Int() != 1 {
		t.Fatalf("bool int() should return a 1 when true")
	}

	if b.String() != "true" {
		t.Fatalf("bool string() should return 'true' when true")
	}

	b.Reset()
	if b.Bool() {
		t.Fatalf("bool reset() should switch bool back to false")
	}

	if b.Int() != 0 {
		t.Fatalf("bool int() should return a 0 when false")
	}

	if b.String() != "false" {
		t.Fatalf("bool string() should return 'false' when false")
	}
}
