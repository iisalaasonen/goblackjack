package game

import (
	"testing"
)

type testPlayer struct {
	Player
	expectedStatus string
	expectedDeck   int
}

func TestHit(t *testing.T) {
	deck := []card{
		{
			value: "3",
			suit:  "clubs",
		},
		{
			value: "10",
			suit:  "hearts",
		},
		{
			value: "6",
			suit:  "clubs",
		},
		{
			value: "Q",
			suit:  "diamonds",
		},
	}

	testPlayers := []testPlayer{
		{
			Player: Player{
				Cards: []card{
					{
						value: "2",
						suit:  "spades",
					},
					{
						value: "8",
						suit:  "clubs",
					},
				},
				Score: 10,
			},
			expectedStatus: "playing",
			expectedDeck:   3,
		},
		{
			Player: Player{
				Cards: []card{
					{
						value: "A",
						suit:  "hearts",
					},
					{
						value: "J",
						suit:  "spades",
					},
				},
				Score: 21,
			},
			expectedStatus: "stand",
			expectedDeck:   3,
		},
		{
			Player: Player{
				Cards: []card{
					{
						value: "10",
						suit:  "clubs",
					},
					{
						value: "7",
						suit:  "diamonds",
					},
				},
				Score: 15,
			},
			expectedStatus: "busted",
			expectedDeck:   2,
		},
	}
	for _, tp := range testPlayers {
		var status string
		deck, status = Hit(&tp.Player, deck)
		if status != tp.expectedStatus {
			t.Errorf("Expected status to be: %v got: %v", tp.expectedStatus, status)
		}
		if len(deck) != tp.expectedDeck {
			t.Errorf("Expected new deck length to be: %d got: %d", tp.expectedDeck, len(deck))
		}
	}
}
