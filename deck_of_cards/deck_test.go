package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card of Ace of Hearts, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Diamonds" {
		t.Errorf("Expected last card of Four of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v cards", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
