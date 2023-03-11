package lc2563

import "sort"

func CountFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)

	var result int64 = 0
	for i := 0; i < len(nums); i++ {
		nextIndex, min, max := i+1, lower-nums[i], upper-nums[i]

		minIndex := sort.Search(len(nums)-nextIndex, func(j int) bool { return nums[j+nextIndex] >= min }) + nextIndex
		maxIndex := sort.Search(len(nums)-minIndex, func(j int) bool { return nums[j+minIndex] > max }) + minIndex

		result += int64(maxIndex - minIndex)
	}

	return result
}
