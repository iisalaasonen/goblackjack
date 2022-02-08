package game

import (
	"testing"
)

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

	testPlayers := []struct {
		Player
		expectedStatus string
		expectedDeck   int
	}{
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

func TestIsSplit(t *testing.T) {
	type testHand struct {
		cards []card
		want  bool
	}
	testHands := []testHand{
		{
			cards: []card{
				{
					value: "10",
					suit:  "clubs",
				},
				{
					value: "10",
					suit:  "hearts",
				},
			},
			want: true,
		},
		{
			cards: []card{
				{
					value: "J",
					suit:  "clubs",
				},
				{
					value: "7",
					suit:  "diamonds",
				},
			},
			want: false,
		},
		{
			cards: []card{
				{
					value: "8",
					suit:  "hearts",
				},
				{
					value: "4",
					suit:  "spades",
				},
			},
			want: false,
		},
	}
	for _, hand := range testHands {
		got, err := IsSplit(hand.cards)
		if err != nil {
			t.Errorf("Expected split to be: %t got error: %s", hand.want, err)
		}
		if got != hand.want {
			t.Errorf("Expected split to be: %t got: %t", hand.want, got)
		}
	}
	errorHand := testHand{
		cards: []card{
			{
				value: "9",
				suit:  "spades",
			},
			{
				value: "9",
				suit:  "hearts",
			},
			{
				value: "4",
				suit:  "clubs",
			},
		},
		want: true,
	}
	_, err := IsSplit(errorHand.cards)
	if err == nil {
		t.Errorf("Split can take only 2 cards. Expected error but got nil")
	}

}

func TestAddCard(t *testing.T) {
	deck := []card{
		{
			value: "5",
			suit:  "clubs",
		},
		{
			value: "K",
			suit:  "hearts",
		},
		{
			value: "A",
			suit:  "clubs",
		},
		{
			value: "4",
			suit:  "spades",
		},
	}
	testPlayers := []struct {
		Player
		wantCardsLen int
		wantDeckLen  int
	}{
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
			wantCardsLen: 3,
			wantDeckLen:  3,
		},
		{
			Player: Player{
				Cards: []card{
					{
						value: "2",
						suit:  "hearts",
					},
					{
						value: "J",
						suit:  "spades",
					},
					{
						value: "5",
						suit:  "spades",
					},
				},
				Score: 21,
			},
			wantCardsLen: 4,
			wantDeckLen:  2,
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
			wantCardsLen: 3,
			wantDeckLen:  1,
		},
	}
	for _, p := range testPlayers {
		deck = AddCard(&p.Player, deck)
		if p.wantCardsLen != len(p.Player.Cards) {
			t.Errorf("Expected player cards length to be: %d got: %d ", p.wantCardsLen, len(p.Player.Cards))
		}
		if p.wantDeckLen != len(deck) {
			t.Errorf("Expected new deck length to be: %d got: %d ", p.wantDeckLen, len(deck))
		}
	}
}

func TestDealerLoop(t *testing.T) {
	deck := []card{
		{
			value: "5",
			suit:  "clubs",
		},
		{
			value: "10",
			suit:  "clubs",
		},
		{
			value: "K",
			suit:  "hearts",
		},
		{
			value: "A",
			suit:  "clubs",
		},
		{
			value: "4",
			suit:  "spades",
		},
	}
	testDealers := []struct {
		Player
		want int
	}{
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
			want: 5,
		},
		{
			Player: Player{
				Cards: []card{
					{
						value: "9",
						suit:  "spades",
					},
					{
						value: "5",
						suit:  "diamonds",
					},
				},
				Score: 11,
			},
			want: 3,
		},
		{
			Player: Player{
				Cards: []card{
					{
						value: "10",
						suit:  "spades",
					},
					{
						value: "Q",
						suit:  "diamonds",
					},
				},
				Score: 20,
			},
			want: 2,
		},
	}
	for _, d := range testDealers {
		DealerLoop(&d.Player, deck)
		if len(d.Player.Cards) != d.want {
			t.Errorf("Expected dealer cards length to be: %d got: %d", d.want, len(d.Player.Cards))
		}
	}
}
