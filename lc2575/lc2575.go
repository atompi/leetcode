package lc2575

func DivisibilityArray(word string, m int) []int {
	ans := make([]int, len(word))
	val := 0

	for i := 0; i < len(word); i++ {
		val = (val*10 + int(word[i]-'0')) % m
		if val%m == 0 {
			ans[i] = 1
		} else {
			ans[i] = 0
		}
	}

	return ans
}
