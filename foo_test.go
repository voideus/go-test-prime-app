package main

import "testing"

func TestFooer(t *testing.T) {
	result := Fooer(4)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}
