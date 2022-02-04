package game

import (
	"math/rand"
	"time"
)

type card struct {
	value string
	suit  string
}

var suites = [4]string{"diamonds", "spades", "clubs", "hearts"}
var values = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

// createDeck creates deck of 52 cards where card is a struct with value and suit
func createDeck() []card {
	deck := make([]card, 0, 52)
	for _, s := range suites {
		for _, v := range values {
			c := card{
				value: v,
				suit:  s,
			}
			deck = append(deck, c)
		}
	}
	return deck
}

// shuffleDeck shuffles the deck using Shuffle from rand library
func shuffleDeck(deck []card) []card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

// Pop removes the last card from the deck and adds that card to player/dealer cards slice
// returns new deck slice which has len-1
func Pop(deck []card) (d []card, c card) {
	l := len(deck)
	c = deck[l-1]
	newDeck := make([]card, 0)
	newDeck = append(newDeck, deck[:l-1]...)
	return newDeck, c
}
