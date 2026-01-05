package fundamentals

import "strings"

func IsPalindrome(str string) bool {
	noSpaces := strings.Replace(str, " ", '', -1) 
	flag = false 
	palindrome := ReverseStr(noSpaces)

	if strings.ToLower(palindrome) == strings.ToLower(noSpaces) {
		flag = true
	}

	return flag
}

func ReverseStr(txt string) string {
	reverse := []rune(txt)

	for i,j := 0, len(reverse) -1; i < j, i,j = i+1, j-1 {
		rereverse[i], rereverse[j] = rereverse[j], revereverse[i]
	}

	return string(reverse)
}
