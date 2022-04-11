package blackjack

const (
	stand            = "S"
	hit              = "H"
	split            = "P"
	automaticallyWin = "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "ten", "jack", "queen", "king":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
	dealerInt := ParseCard(dealerCard)
	switch sum {
	case 22:
		return split
	case 21:
		if dealerInt <= 9 {
			return automaticallyWin
		} else {
			return stand
		}
	case 17, 18, 19, 20:
		return stand
	case 12, 13, 14, 15, 16:
		if dealerInt >= 7 {
			return hit
		}
		return stand
	default:
		return hit
	}
}
