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
			var splitChoice string
			var split bool
			var splitPlayer game.Player
			isSplit := game.IsSplit(player.Cards)
			if isSplit {
				fmt.Println("YOUR HAND: ", game.PrintHand(player.Cards))
				fmt.Println("WANT TO SPLIT (Y OR N): ")
				fmt.Scan(&splitChoice)
				if splitChoice != "Y" && splitChoice != "y" {
					split = false
				} else {
					split = true
				}
			}
			if split {
				var status1, status2 string
				splitPlayer.Cards = append(splitPlayer.Cards, player.Cards[1])
				player.Cards = player.Cards[:1]
				fmt.Println("FIRST SPLIT CARD:")
				player.Score = game.CalculateScore(player.Cards)
				status1, deck = game.Playing(&player, deck)
				fmt.Println("SECOND SPLIT CARD:")
				splitPlayer.Score = game.CalculateScore(splitPlayer.Cards)
				status2, deck = game.Playing(&splitPlayer, deck)
				fmt.Println("FIRST SPLIT CARD RESULT:")
				game.ShowResult(status1, &dealer, &player, deck)
				fmt.Println("SECOND SPLIT CARD RESULT:")
				game.ShowResult(status2, &dealer, &splitPlayer, deck)
			} else {
				status, deck := game.Playing(&player, deck)
				game.ShowResult(status, &dealer, &player, deck)
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
