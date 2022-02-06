package game

import (
	"testing"
)

func TestCreateDeck(t *testing.T) {
	deck := createDeck()
	if len(deck) != 52 {
		t.Error("Expected length of the deck to be ", 52)
	}
	numberOfAces := 0
	for _, card := range deck {
		value := card.value
		if value == "A" {
			numberOfAces++
		}
	}
	if numberOfAces != 4 {
		t.Errorf("Expected %d aces got %d", 4, numberOfAces)
	}
}

func TestPop(t *testing.T) {
	deck := createDeck()
	beforeLen := len(deck)
	afterLen := beforeLen - 1
	poppedCard := deck[beforeLen-1]
	newDeck, c := Pop(deck)
	if len(newDeck) != afterLen {
		t.Error("")
	}
	if poppedCard.value != c.value && poppedCard.suit != c.suit {
		t.Error("Cards are not the same")
	}

}
