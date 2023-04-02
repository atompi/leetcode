package lc2574

func LeftRigthDifference(nums []int) []int {
	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))
	ans := make([]int, len(nums))
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		if i == 0 {
			leftSum[i] = 0
			continue
		}
		if j == len(nums)-1 {
			rightSum[j] = 0
			continue
		}
		leftSum[i] = leftSum[i-1] + nums[i-1]
		rightSum[j] = rightSum[j+1] + nums[j+1]
	}
	for i := 0; i < len(ans); i++ {
		ans[i] = abs(leftSum[i] - rightSum[i])
	}
	return ans
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
