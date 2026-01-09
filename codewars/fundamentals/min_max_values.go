package fundamentals

func MinMax(arr []int) [2]int {
	min, max := arr[0], arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}

	return [2]int{min, max}
}
