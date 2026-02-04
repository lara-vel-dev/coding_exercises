package fundamentals

func BinToDec(bin string) int {
	num := 0

	for i := 0; i < len(bin); i++ {
		num = num*2 + int(bin[i]-'0')
	}

	return num
}
