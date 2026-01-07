package fundamentals

import "strconv"

func Digitize(n int) (digits []int) {
	runes := []rune(strconv.Itoa(n))

	for i := len(runes) - 1; i >= 0; i-- {
		digits = append(digits, int(runes[i]-'0'))
	}

	return
}
