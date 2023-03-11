package lc724

func PivotIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	sums := make([][2]int, len(nums))

	sum := 0
	for _, v := range nums {
		sum += v
	}
	sums[0][0] = 0
	sums[0][1] = sum - nums[0]
	if sums[0][0] == sums[0][1] {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		sums[i][0] = sums[i-1][0] + nums[i-1]
		sums[i][1] = sums[i-1][1] - nums[i]
		if sums[i][0] == sums[i][1] {
			return i
		}
	}
	return -1
}
