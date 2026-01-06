package fundamentals

func Century(year int) int {
	if year%100 == 0 {
		return int(year / 100)
	}

	return int(year/100) + 1
}
