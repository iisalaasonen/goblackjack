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
		var choice string
		for i := 0; i < 2; i++ {
			deck = game.Pop(deck, &player)
			deck = game.Pop(deck, &dealer)
		}
		game.CalculateScore(&player)
		game.CalculateScore(&dealer)
		fmt.Println("DEALER HAS ", dealer.Cards[0])
		if player.Score == 21 && dealer.Score == 21 {
			fmt.Println("BOTH HAVE BLACKJACK!")
		} else if player.Score == 21 {
			fmt.Println("YOU HAVE BLACKJACK!")
		} else if dealer.Score == 21 {
			fmt.Println("YOU LOSE!")
		} else {
			status := game.PlayerLoop(&player, deck)
			fmt.Println("YOUR HAND ", game.PrintHand(player.Cards))
			fmt.Println("YOUR SCORE ", player.Score)
			switch status {
			case "busted":
				fmt.Println("YOU LOSE")
			case "stand":
				game.DealerLoop(&dealer, deck)
				fmt.Println("DEALER HAND ", game.PrintHand(dealer.Cards))
				fmt.Println("DEALER SCORE ", dealer.Score)
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
		fmt.Println("PLAY AGAIN (P) OR QUIT (Q): ")
		fmt.Scan(&choice)
		if choice == "Q" || choice == "q" {
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
