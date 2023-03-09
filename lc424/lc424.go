package lc424

func CharacterReplacement(s string, k int) int {
	charMap := map[byte]int{}
	maxLen := 0
	maxRepeat := byte(0)
	left := 0
	right := 0

	for right < len(s) {
		char := s[right]
		charMap[char]++
		count := charMap[char]

		if maxRepeat == 0 || charMap[maxRepeat] < count {
			maxRepeat = char
		}

		if right-left+1-charMap[maxRepeat] > k {
			charMap[s[left]]--
			left++
		}

		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
		right++
	}

	return maxLen
}
