package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten":
		value = 10
	case "jack":
		value = 10
	case "queen":
		value = 10
	case "king":
		value = 10
	case "other":
		value = 0
	}

	return value
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	cardValue := ParseCard(card1) + ParseCard(card2)

	return cardValue == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	var decision string

	switch {
	case isBlackjack && dealerScore < 10:
		decision = "W"
	case isBlackjack && dealerScore >= 10:
		decision = "S"
	case !isBlackjack:
		decision = "P"
	}

	return decision
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	var decision string

	switch {
	case handScore > 20:
		decision = LargeHand(true, dealerScore)

	case handScore == 22:
		decision = LargeHand(false, dealerScore)

	case handScore < 21 && handScore >= 17:
		decision = "S"

	case handScore <= 11:
		decision = "H"

	case handScore <= 16 && handScore >= 12:
		if dealerScore >= 7 {
			decision = "H"
		} else {
			decision = "S"
		}
	}

	return decision
}
