package fundamentals

func PositiveSum(numbers []int) (total int) {
	for _, num := range numbers {
		if num > 0 {
			total += num
		}
	}

	return
}
