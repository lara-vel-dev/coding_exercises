package fundamentals

func FindShort(s string) int {
	runes := []rune(s)
	minVal := len(runes)
	count := 0

	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			if count > 0 && count < minVal {
				minVal = count
			}
			count = 0
			continue
		}
		count++
	}

	return minVal
}
