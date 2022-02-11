package game

import (
	"os"
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

func TestPlaying(t *testing.T) {
	var gotStatus string
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
		*Player
		input          string
		expectedStatus string
		expectedDeck   int
	}{
		{Player: &Player{
			Cards: []card{
				{
					value: "6",
					suit:  "spades",
				},
				{
					value: "8",
					suit:  "clubs",
				},
			},
			Score: 14,
		},
			input:          "s",
			expectedStatus: "stand",
			expectedDeck:   4},
		{Player: &Player{
			Cards: []card{
				{
					value: "4",
					suit:  "spades",
				},
				{
					value: "6",
					suit:  "clubs",
				},
				{
					value: "Q",
					suit:  "clubs",
				},
			},
			Score: 20,
		},
			input:          "h",
			expectedStatus: "busted",
			expectedDeck:   3},
	}
	for _, p := range testPlayers {
		input := []byte(p.input)
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}

		_, err = w.Write(input)
		if err != nil {
			t.Error(err)
		}
		w.Close()

		stdin := os.Stdin
		// Restore stdin right after the test.
		defer func() { os.Stdin = stdin }()
		os.Stdin = r
		gotStatus, deck = Playing(p.Player, deck)
		if gotStatus != p.expectedStatus {
			t.Errorf("Expected %s got: %s", p.expectedStatus, gotStatus)
		}
		if len(deck) != p.expectedDeck {
			t.Errorf("Expected new deck length to be %d got %d: ", p.expectedDeck, len(deck))
		}
	}
	hitPlayer := struct {
		*Player
		expectedStatus string
		expectedDeck   int
	}{
		Player: &Player{
			Cards: []card{
				{
					value: "3",
					suit:  "hearts",
				},
				{
					value: "2",
					suit:  "clubs",
				},
			},
			Score: 5,
		},
		expectedStatus: "stand",
		expectedDeck:   0,
	}
	input := []byte("h")
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(input)
	if err != nil {
		t.Error(err)
	}
	w.Close()

	stdin := os.Stdin
	// Restore stdin right after the test.
	defer func() { os.Stdin = stdin }()
	os.Stdin = r
	gotStatus, deck = Playing(hitPlayer.Player, deck)
	if gotStatus != "stand" {
		t.Error("Expected stand got: ", gotStatus)
	}
	if len(deck) != 1 {
		t.Errorf("Expected new deck length to be 1 got: %d", len(deck))
	}
}
