package game

type Player struct {
	Cards []card
	Score int
}

func parseCard(c card) int {
	var value int
	switch c.value {
	case "A":
		value = 1
	case "10", "J", "Q", "K":
		value = 10
	case "2":
		value = 2
	case "3":
		value = 3
	case "4":
		value = 4
	case "5":
		value = 5
	case "6":
		value = 6
	case "7":
		value = 7
	case "8":
		value = 8
	case "9":
		value = 9
	default:
		value = 0
	}
	return value
}

func CalculateScore(player *Player) {
	hasAce := false
	score := 0
	for _, card := range player.Cards {
		if card.value == "A" {
			hasAce = true
		}
		cardValue := parseCard(card)
		score += cardValue

	}
	if hasAce {
		if score <= 11 {
			score += 10
		}
	}
	player.Score = score
}

func PrintHand(cards []card) string {
	var hand string
	for _, c := range cards {
		hand += c.value + " " + c.suit + " "
	}
	return hand
}
