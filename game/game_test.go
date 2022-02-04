package game

import "testing"

func TestCreateDeck(t *testing.T) {
	deck := createDeck()
	if len(deck) != 52 {
		t.Error("Expected length of the deck to be ", 52)
	}
}
