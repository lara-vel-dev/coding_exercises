package fundamentals

func QuarterOf(month int) int {
	/*
		(x + (k-1))/k
		x = month
		k = 3
	*/
	return (month + 2) / 3
}
