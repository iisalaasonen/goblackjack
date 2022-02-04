package main

import (
	"fmt"
	"os"

	"github.com/iisalaasonen/goblackjack/game"
)

func main() {
	printBJ()
	for {
		deck := game.InitializeDeck()
		var player game.Player
		var dealer game.Player
		for i := 0; i < 2; i++ {
			deck = game.AddCard(&player, deck)
			deck = game.AddCard(&dealer, deck)
		}
		player.Score = game.CalculateScore(player.Cards)
		dealer.Score = game.CalculateScore(dealer.Cards)
		fmt.Println("DEALER HAS ", dealer.Cards[0])
		if player.Score == 21 && dealer.Score == 21 {
			fmt.Println("BOTH HAVE BLACKJACK!")
		} else if player.Score == 21 {
			fmt.Println("YOUR HAND: ", game.PrintHand(player.Cards))
			fmt.Println("YOU GOT BLACKJACK!")
		} else if dealer.Score == 21 {
			fmt.Println("DEALER HAND: ", game.PrintHand(dealer.Cards))
			fmt.Println("DEALER GOT BLACKJACK! YOU LOSE!")
		} else {
			var choice string
			status := "playing"
			isSplit := game.IsSplit(player.Cards)
			fmt.Println("SPLIT ", isSplit)
			for status == "playing" {
				fmt.Println("YOUR HAND: ", game.PrintHand(player.Cards))
				fmt.Println("YOUR SCORE: ", player.Score)
				fmt.Println("HIT (H) OR STAND (S): ")
				fmt.Scan(&choice)
				switch choice {
				case "H", "h":
					deck, status = game.Hit(&player, deck)
				case "S", "s":
					status = "stand"
				}
			}
			switch status {
			case "busted":
				fmt.Println("BUSTED WITH SCORE: ", player.Score)
				fmt.Println("YOU LOSE")
			case "stand":
				game.DealerLoop(&dealer, deck)
				fmt.Println("DEALER HAND: ", game.PrintHand(dealer.Cards))
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
		var again string
		fmt.Println("PLAY AGAIN (P) OR QUIT (Q): ")
		fmt.Scan(&again)
		if again == "Q" || again == "q" {
			os.Exit(1)
		}

	}
}

func printBJ() {
	fmt.Println("#######        ###########")
	fmt.Println("##    ##            ##    ")
	fmt.Println("## ##  ##           ##    ")
	fmt.Println("## ##  ##           ##    ")
	fmt.Println("##    ##            ##    ")
	fmt.Println("#######             ##    ")
	fmt.Println("##    ##   ##       ##    ")
	fmt.Println("## ##  ##  ##       ##    ")
	fmt.Println("## ##  ##  ##       ##    ")
	fmt.Println("##    ##    ##     ##     ")
	fmt.Println("#######       #####       ")
}
