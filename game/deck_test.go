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
	var deck []card
	var poppedCard card
	deck = append(deck, card{
		value: "5",
		suit:  "clubs",
	},
		card{
			value: "10",
			suit:  "clubs",
		},
		card{
			value: "K",
			suit:  "hearts",
		},
		card{
			value: "A",
			suit:  "clubs",
		},
		card{
			value: "4",
			suit:  "spades",
		})
	popCards := []struct {
		wantCard card
		wantLen  int
	}{
		{
			wantCard: card{
				value: "4",
				suit:  "spades",
			},
			wantLen: 4,
		},
		{
			wantCard: card{
				value: "A",
				suit:  "clubs",
			},
			wantLen: 3,
		},
	}
	for _, c := range popCards {
		deck, poppedCard = Pop(deck)
		if c.wantCard != poppedCard {
			t.Errorf("Expected popped card to be: %+v got: %+v", c.wantCard, poppedCard)
		}
		if len(deck) != c.wantLen {
			t.Errorf("Expected new deck length to be: %d got: %d", c.wantLen, len(deck))
		}

	}

}

/*
func TestShuffle(t *testing.T) {

}
*/
