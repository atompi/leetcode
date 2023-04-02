package lc509

func Fib(n int) int {
	if n <= 1 {
		return n
	}

	a, b, c := 0, 1, 0
	for n > 1 {
		c = a + b
		a, b = b, c
		n--
	}
	return c
}
