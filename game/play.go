/*
Package game implements logic of a blackjack game including
deck, player, dealer, scoring and play options
*/
package game

// InitializeDeck creates new 52 card deck and shuffles it
func InitializeDeck() []card {
	newDeck := createDeck()
	shuffleDeck(newDeck)
	return newDeck
}

//Hit adds card to player cards and checks the score
//returns player game status, player can bust or keep playing
func Hit(player *Player, deck []card) ([]card, string) {
	if player.Score == 21 {
		return deck, "stand"
	}
	newDeck := AddCard(player, deck)
	player.Score = CalculateScore(player.Cards)
	if player.Score > 21 {
		return newDeck, "busted"
	}
	return newDeck, "playing"
}

//IsSplits checks if player has an option to split two cards
//split is possible if two cards are same and card value is over 9
func IsSplit(cards []card) bool {
	if cards[0].value == cards[1].value {
		score := CalculateScore(cards)
		if score >= 18 {
			return true
		}
	}
	return false
}

//AddCard first pops card from deck and adds that card to the player cards
//returns deck with minus 1 length
func AddCard(player *Player, deck []card) []card {
	d, c := Pop(deck)
	player.Cards = append(player.Cards, c)
	return d
}

// DealerLoop gives dealer score
func DealerLoop(dealer *Player, deck []card) {
	// dealer hits until at least 17 is reached
	for dealer.Score < 17 {
		deck = AddCard(dealer, deck)
		dealer.Score = CalculateScore(dealer.Cards)
	}
}
