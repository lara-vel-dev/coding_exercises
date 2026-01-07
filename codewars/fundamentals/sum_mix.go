package fundamentals

import "strconv"

func SumMix(arr []any) (total int) {
	for _, v := range arr {
		switch val := v.(type) {
		case int:
			total += val
		case string:
			n, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			total += n
		}
	}
	return
}
