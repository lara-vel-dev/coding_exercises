package fundamentals

func RemoveChar(word string) string {
	runes := []rune(word)
	newStr := []rune{}

	for i := 1; i < len(runes)-1; i++ {
		newStr = append(newStr, runes[i])
	}

	return string(newStr)
}
