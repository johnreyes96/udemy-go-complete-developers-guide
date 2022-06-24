package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestToString(t *testing.T) {
	d := newDeck()
	deckJoinWithCommas := d.toString()
	deckSplit := strings.Split(deckJoinWithCommas, ",")

	if deckSplit[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", deckSplit[0])
	}
	if deckSplit[len(deckSplit)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", deckSplit[len(deckSplit)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	err := os.Remove("_decktesting")
	if err != nil {
		return
	}

	deck := newDeck()
	err = deck.saveToFile("_decktesting")
	if err != nil {
		return
	}

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	err = os.Remove("_decktesting")
	if err != nil {
		return
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	d.shuffle()

	if len(d) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(d))
	}
}
