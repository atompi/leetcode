package lc70

func ClimbStairs(n int) int {
	if n <= 1 {
		return n
	}
	a, b, c := 1, 2, 0
	for n > 2 {
		c = a + b
		a, b = b, c
		n--
	}
	return b
}
