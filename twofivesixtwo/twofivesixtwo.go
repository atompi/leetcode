package twofivesixtwo

import (
	"strconv"
)

func FindTheArrayConcVal(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}

	var result int64 = 0
	for i := 0; i < len(nums)/2; i++ {
		n1 := strconv.Itoa(nums[i])
		n2 := strconv.Itoa(nums[len(nums)-1-i])
		tmp, _ := strconv.Atoi(n1 + n2)
		result += int64(tmp)
	}
	if len(nums)%2 != 0 {
		result += int64(nums[len(nums)/2])
	}
	return result
}
