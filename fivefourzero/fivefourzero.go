package fivefourzero

func SingleNonDuplicate(nums []int) int {
	ans := nums[0]
	low, high := 0, len(nums)-1
	for low <= high {
		middle := (low + high) / 2
		if middle == 0 || middle == len(nums)-1 {
			ans = nums[middle]
			break
		}
		if nums[middle] != nums[middle-1] && nums[middle] != nums[middle+1] {
			ans = nums[middle]
			break
		}
		if middle%2 != 0 {
			if nums[middle] == nums[middle+1] {
				high = middle - 1
			} else {
				low = middle + 1
			}
		} else {
			if nums[middle] == nums[middle+1] {
				low = middle + 1
			} else {
				high = middle - 1
			}
		}
	}
	return ans
}
