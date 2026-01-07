package fundamentals

func CountSheeps(sheep []bool) (total int) {
	for _, count := range sheep {
		if count {
			total++
		}
	}

	return
}
