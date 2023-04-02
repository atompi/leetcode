package lc35

func SearchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
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
	return low
}
