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
		// you can also validate str > m (ASCII value is evaluated)
		if !strings.Contains(letters, string(str)) {
			count++
		}
	}

	return fmt.Sprintf("%d/%d", count, len(runeStr))

}
