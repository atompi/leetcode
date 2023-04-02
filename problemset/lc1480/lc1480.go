package lc1480

func RunningSum(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	sums := make([]int, len(nums))
	for k, v := range nums {
		if k == 0 {
			sums[k] = nums[k]
		} else {
			sums[k] = sums[k-1] + v
		}
	}
	return sums
}
