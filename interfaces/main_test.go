package main

import "testing"

func TestGetGreetingSpanish(t *testing.T) {
	sb := spanishBot{}

	greeting := sb.getGreeting()

	if greeting != "Hola!" {
		t.Errorf("Expected Hola!, but got %v", greeting)
	}
}

func TestGetGreetingEnglish(t *testing.T) {
	sb := englishBot{}

	greeting := sb.getGreeting()

	if greeting != "Hi There!" {
		t.Errorf("Expected Hi There!, but got %v", greeting)
	}
}
