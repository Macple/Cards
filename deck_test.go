package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Wrong number of cards in the deck. Should be 16. Is: %d", len(d))
	}

	if d[0] != "As Kier" {
		t.Errorf("Expected 'As Kier' on the first position, but got '%v'", d[0])
	}

	if d[len(d)-1] != "Czwórka Pik" {
		t.Errorf("Expected 'Czwórka Pik' on the last position, but got '%v'", d[len(d)-1])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected number of cards in the deck is 16, but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
