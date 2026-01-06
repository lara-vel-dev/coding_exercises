package fundamentals

func CalculateYears(years int) (result [3]int) {
	cat, dog := 15, 15
	yrs := years - 2

	switch years {
	case 1:
		result = [3]int{years, cat, dog}
	case 2:
		result = [3]int{years, cat + 9, dog + 9}
	default:
		cat, dog = cat+9, dog+9
		result = [3]int{years, cat + (4 * yrs), dog + (5 * yrs)}
	}

	return
}
