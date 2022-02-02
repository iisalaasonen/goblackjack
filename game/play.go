package game

import (
	"fmt"
)

func InitializeDeck() []card {
	newDeck := createDeck()
	shuffleDeck(newDeck)
	return newDeck
}

func PlayerLoop(player *Player, deck []card) string {
	var status string
	if player.Cards[0].value == player.Cards[1].value {
		fmt.Println("SPLIT OPTION")
	}
	for {
		fmt.Println("YOUR HAND ", PrintHand(player.Cards))
		fmt.Println("YOUR SCORE ", player.Score)
		var option string
		fmt.Println("HIT (H) OR STAND (S): ")
		fmt.Scan(&option)
		switch option {
		case "H", "h":
			deck = Pop(deck, player)
			CalculateScore(player)
			if player.Score > 21 {
				status = "busted"
				return status
			}
		case "S", "s":
			status = "stand"
			return status
		}
	}
}

func DealerLoop(dealer *Player, deck []card) {
	for dealer.Score < 17 {
		deck = Pop(deck, dealer)
		CalculateScore(dealer)
	}
}
