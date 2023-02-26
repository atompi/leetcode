package twofivesevenfour

func LeftRigthDifference(nums []int) []int {
	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))
	ans := make([]int, len(nums))
	j := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			leftSum[i] = 0
			j--
			continue
		}
		if j == len(nums)-1 {
			rightSum[j] = 0
			j--
			continue
		}
		leftSum[i] = leftSum[i-1] + nums[i-1]
		rightSum[j] = rightSum[j+1] + nums[j+1]
		j--
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
