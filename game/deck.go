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

func shuffleDeck(deck []card) []card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

func Pop(deck []card, player *Player) (d []card) {
	l := len(deck)
	c := deck[l-1]
	player.Cards = append(player.Cards, c)
	newDeck := make([]card, 0)
	newDeck = append(newDeck, deck[:l-1]...)
	return newDeck
}
