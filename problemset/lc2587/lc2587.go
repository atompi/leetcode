package lc2587

import (
	"sort"
)

func MaxScore(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	prefixSum := 0
	score := 0
	for _, num := range nums {
		prefixSum += num
		if prefixSum > 0 {
			score++
		}
	}
	return score
}
