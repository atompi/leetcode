package ninezerofour

import "github.com/atompi/leetcode/utils/max"

func TotalFruit(fruits []int) int {
	left, right, counter, res, freq := 0, 0, 1, 1, map[int]int{}
	freq[fruits[0]]++
	for left < len(fruits) {
		if right+1 < len(fruits) && ((counter > 0 && fruits[right+1] != fruits[left]) || (fruits[right+1] == fruits[left] || freq[fruits[right+1]] > 0)) {
			if counter > 0 && fruits[right+1] != fruits[left] {
				counter--
			}
			right++
			freq[fruits[right]]++
		} else {
			if counter == 0 || (counter > 0 && right == len(fruits)-1) {
				res = max.MaxOfTwoInt(res, right-left+1)
			}
			freq[fruits[left]]--
			if freq[fruits[left]] == 0 {
				counter++
			}
			left++
		}
	}
	return res
}
