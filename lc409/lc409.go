package lc409

func LongestPalindrome(s string) int {
	counter := make(map[rune]int)
	for _, r := range s {
		counter[r]++
	}
	result := 0
	for _, v := range counter {
		result += (v / 2) * 2
		if result%2 == 0 && v%2 == 1 {
			result++
		}
	}
	return result
}
