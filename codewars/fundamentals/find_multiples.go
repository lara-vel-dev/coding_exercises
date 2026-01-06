package fundamentals

func FindMultiples(integer int, limit int) (multiples []int) {
	for i := integer; i <= limit; i++ {
		if i%integer == 0 {
			multiples = append(multiples, i)
		}
	}

	return
}
