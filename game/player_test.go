package game

import (
	"testing"
)

func TestParseCard(t *testing.T) {
	testCards := []struct {
		card
		want int
	}{
		{
			card: card{
				value: "A",
				suit:  "clubs",
			},
			want: 1,
		},
		{
			card: card{value: "2",
				suit: "spades"},
			want: 2,
		},
		{
			card: card{value: "8",
				suit: "clubs"},
			want: 8,
		},
		{
			card: card{value: "10",
				suit: "hearts"},
			want: 10},
		{
			card: card{value: "5",
				suit: "diamonds"},
			want: 5},
		{
			card: card{value: "-7",
				suit: "diamonds"},
			want: 0},
		{
			card: card{value: "14",
				suit: "hearts"},
			want: 0},
	}
	for _, c := range testCards {
		got := parseCard(c.card)
		if got != c.want {
			t.Errorf("Expected to get value: %d got: %d", c.want, got)
		}
	}
}

func TestCalculateScore(t *testing.T) {
	testHands := []struct {
		cards []card
		want  int
	}{
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
				{
					value: "4",
					suit:  "clubs",
				},
			},
			want: 24,
		},
		{
			cards: []card{
				{
					value: "4",
					suit:  "hearts",
				},
				{
					value: "9",
					suit:  "hearts",
				},
				{
					value: "A",
					suit:  "clubs",
				},
			},
			want: 14,
		},
		{
			cards: []card{
				{
					value: "5",
					suit:  "diamonds",
				},
				{
					value: "6",
					suit:  "hearts",
				},
				{
					value: "2",
					suit:  "clubs",
				},
			},
			want: 13,
		},
		{
			cards: []card{
				{
					value: "A",
					suit:  "diamonds",
				},
				{
					value: "6",
					suit:  "hearts",
				},
				{
					value: "A",
					suit:  "clubs",
				},
			},
			want: 18,
		},
		{
			cards: []card{
				{
					value: "A",
					suit:  "diamonds",
				},
				{
					value: "6",
					suit:  "hearts",
				},
			},
			want: 17,
		},
		{
			cards: []card{
				{
					value: "-5",
					suit:  "diamonds",
				},
				{
					value: "6",
					suit:  "hearts",
				},
			},
			want: 6,
		},
	}
	for _, c := range testHands {
		got := CalculateScore(c.cards)
		if got != c.want {
			t.Errorf("Expected score to be: %d got: %d", c.want, got)
		}
	}
}
