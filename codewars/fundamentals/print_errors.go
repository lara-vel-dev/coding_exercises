package fundamentals

import (
	"fmt"
	"strings"
)

func PrinterError(s string) string {
	letters := "abcdefghijklm"
	runeStr := []rune(s)
	count := 0

	for _, str := range runeStr {
		// validates the rest of the letters after m
		if !strings.Contains(letters, string(str)) {
			count++
		}
	}

	return fmt.Sprintf("%d/%d", count, len(runeStr))

}
