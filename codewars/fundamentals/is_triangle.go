package fundamentals

func IsTriangle(a, b, c int) bool {
	// Triangle Inequality Theorem
	return a+b > c && a+c > b && b+c > a
}
