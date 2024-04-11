package main

import (
	"os"
	"strings"
	"testing"
)

func TestGivenACardsWhenInvokeNewDeckThenMustReturn16Cards(t *testing.T) {
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

func TestGivenNewDeckWhenDeckToStringThenMustReturnCardsInString(t *testing.T) {
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

func TestGivenSaveToDeckWhenNewDeckFromFileThenMustReturnLoadedDeck(t *testing.T) {
	_ = os.Remove("_decktesting")
	deck := newDeck()

	err := deck.saveToFile("_decktesting")

	if err != nil {
		return
	}

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	_ = os.Remove("_decktesting")
}

func TestGivenNewDeckWhenShuffleCardsThenMustReturnCardsShuffled(t *testing.T) {
	d := newDeck()

	d.shuffle()

	if len(d) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(d))
	}
}

func TestGivenNewDeckWhenDealCardsThenMustReturn2DecksDealed(t *testing.T) {
	d := newDeck()

	d1, d2 := deal(d, 5)

	if len(d1) != 5 {
		t.Errorf("Expected 5 cards in deck, but got %v", len(d1))
	}
	if len(d2) != 11 {
		t.Errorf("Expected 11 cards in deck, but got %v", len(d2))
	}
}
