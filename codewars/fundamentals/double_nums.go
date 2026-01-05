package fundamentals

func Maps(x []int) (double []int) {

	for _, num := range x {
		double = append(double, num*2)
	}

	return
}
