package fundamentals

func StringToNumber(str string) int {
	sign := 1
	start := 0
	num := 0

	if str[0] == '-' {
		sign = -1
		start = 1
	}

	for i := start; i < len(str); i++ {
		num = num*10 + int(str[i]-'0')
	}

	return num * sign
}
