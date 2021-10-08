package blackjack

var cards = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cards[card]
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	v1, v2 := cards[card1], cards[card2]
	return (v1 + v2) == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack && dealerScore < 10 {
		return "W"
	}
	if !isBlackjack {
		return "P"
	}
	return "S"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch {
	case handScore <= 11:
		return "H"
	case handScore < 17 && dealerScore >= 7:
		return "H"
	default:
		return "S"
	}
}
