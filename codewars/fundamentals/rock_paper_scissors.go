package fundamentals

import "strings"

func Rsp(p1, p2 string) string {
	p1 = strings.ToLower(p1)
	p2 = strings.ToLower(p2)

	if p1 == p2 {
		return "Draw!"
	}

	wins := map[string]string{
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	}

	if wins[p1] == p2 {
		return "Player 1 wins!"
	}

	return "Player 2 wins!"
}
