package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	link := "https://google.com.co"
	c := make(chan string)

	go checkLink(link, c)

	url := <-c
	if url != "https://google.com.co" {
		t.Errorf("Expected https://google.com.co, but got %v", url)
	}
}
