package lc231

func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// lowbit
func IsPowerOfTwo2(n int) bool {
	return n > 0 && n == (n&(-n))
}
