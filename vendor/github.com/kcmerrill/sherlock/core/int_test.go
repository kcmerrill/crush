package core

import "testing"

func TestIntParam(t *testing.T) {
	i := NewInt()

	// set to 1000
	i.Set(1000)

	if val := i.String(); val != "1000" {
		t.Fatalf("Expected '1000', Actual: %s", val)
	}

	// test reset
	i.Reset()

	if empty := i.Int(); empty != 0 {
		t.Fatalf("Expected Int Reset() to reset the value to 0")
	}

	// verify created is not empty
	if i.Created().IsZero() {
		t.Fatalf("Expected NewInt() to init Int.created to time.Time")
	}

	// verify lastmodified is not empty
	if i.LastModified().IsZero() {
		t.Fatalf("Expected NewInt() to init Int.modified to time.Time")
	}

	// test many/read writes
	for x := 0; x <= 300; x++ {
		go func() {
			i.Add(x)
		}()
	}
}
