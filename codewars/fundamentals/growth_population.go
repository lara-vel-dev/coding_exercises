package fundamentals

func NbYear(p0 int, aug int, p int, percent float64) (years int) {
	for p0 < p {
		p0 += int(float64(p0)*(percent/100)) + aug
		years++
	}

	return

}
