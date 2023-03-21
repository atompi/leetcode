package lc2592

import (
	"sort"
)

func MaximizeGreatness(nums []int) int {
	sort.Ints(nums)
	i := 0
	for _, x := range nums {
		if x > nums[i] {
			i++
		}
	}
	return i
}
