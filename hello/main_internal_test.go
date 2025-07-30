package main

import "testing"

func TestGreet_English(t *testing.T) {
	lang := language("en")
	want := "Hello World"
	got := greet(lang) //#2
	if got != want {   //#3
		t.Errorf("expected: %q, got :%q", want, got)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")
	want := "Bonjour le monde"
	got := greet(lang) //#2
	if got != want {   //#3
		t.Errorf("expected: %q, got :%q", want, got)
	}
}

func TestGreet_Spanish(t *testing.T) {
	lang := language("es")
	want := ":=)"
	got := greet(lang) //#2
	if got != want {   //#3
		t.Errorf("expected: %q, got :%q", want, got)
	}
}
