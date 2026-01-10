package fundamentals

import "strconv"

func NbDig(n int, d int) int {
	count := 0
	target := rune('0' + d)

	for i := 0; i <= n; i++ {
		square := i * i
		for _, num := range strconv.Itoa(square) {
			if num == target {
				count++
			}
		}
	}

	return count
}
