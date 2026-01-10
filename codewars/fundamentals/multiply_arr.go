package fundamentals

func Grow(arr []int) int {
	total := 1

	for _, num := range arr {
		total *= num
	}

	return total
}
