package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Wrong number of cards in the deck. Should be 16. Is: %d", len(d))
	}
}
