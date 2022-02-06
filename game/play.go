/*
Package game implements logic of a blackjack game including
deck, player, dealer, scoring and play options
*/
package game

import (
	"fmt"
)

// InitializeDeck creates new 52 card deck and shuffles it
func InitializeDeck() []card {
	newDeck := createDeck()
	shuffleDeck(newDeck)
	return newDeck
}

//Hit adds card to player cards and checks the score
//returns player game status, player can bust or keep playing
func Hit(player *Player, deck []card) ([]card, string) {
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

func Playing(player *Player, deck []card) (string, []card) {
	var choice string
	status := "playing"
	for status == "playing" {
		if player.Score == 21 {
			return "stand", deck
		}
		fmt.Println("YOUR HAND: ", PrintHand(player.Cards))
		fmt.Println("YOUR SCORE: ", player.Score)
		fmt.Println("HIT (H) OR STAND (S): ")
		fmt.Scan(&choice)
		switch choice {
		case "H", "h":
			deck, status = Hit(player, deck)
		case "S", "s":
			status = "stand"
		}
	}
	return status, deck
}

func ShowResult(status string, dealer *Player, player *Player, deck []card) {
	switch status {
	case "busted":
		fmt.Println("BUSTED WITH SCORE: ", player.Score)
		fmt.Println("YOU LOSE")
	case "stand":
		DealerLoop(dealer, deck)
		fmt.Println("DEALER HAND: ", PrintHand(dealer.Cards))
		fmt.Println("DEALER SCORE: ", dealer.Score)
		if dealer.Score > 21 {
			fmt.Println("YOU WON!")
		} else if dealer.Score == player.Score {
			fmt.Println("PUSH!")
		} else if dealer.Score > player.Score {
			fmt.Println("YOU LOSE!")
		} else {
			fmt.Println("YOU WON!")
		}
	}
}
