package lc704

func Search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	low := 0
	high := len(nums) - 1
	for low <= high {
		middle := (low + high) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return -1
}
