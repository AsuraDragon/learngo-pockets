package main

import "testing"

func TestGreet(t *testing.T) {
	want := "Hello World" //#1
	got := greet()        //#2
	if got != want {      //#3
		t.Errorf("expected: %q, got :%q", want, got)
	}
}
