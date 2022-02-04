package game

type Player struct {
	Cards []card
	Score int
}

// parseCard converts card string value to int
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

// CalculateScore calculates player/dealer score of their current hand/cards
func CalculateScore(cards []card) int {
	hasAce := false
	score := 0
	for _, card := range cards {
		if card.value == "A" {
			hasAce = true
		}
		cardValue := parseCard(card)
		score += cardValue

	}
	// ace has default value of 1 (only one ace can have value 11)
	// if possible add 10 to make one ace as value 11
	if hasAce {
		if score <= 11 {
			score += 10
		}
	}
	return score
}

// PrintHand takes cards slice and returns hand values as string
func PrintHand(cards []card) string {
	var hand string
	for _, c := range cards {
		hand += c.value + " " + c.suit + " "
	}
	return hand
}
