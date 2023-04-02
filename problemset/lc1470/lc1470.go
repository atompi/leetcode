package lc1470

func Shuffle(nums []int, n int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums)/2; i++ {
		for j := 0; j*n < len(nums); j++ {
			result = append(result, nums[j*n+i])
		}
	}
	return result
}
