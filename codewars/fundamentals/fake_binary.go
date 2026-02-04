package fundamentals

func FakeBin(x string) string {
	bin := ""

	for _, digit := range x {
		if digit < '5' {
			bin += "0"
		} else {
			bin += "1"
		}
	}

	return bin
}
