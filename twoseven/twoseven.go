package twoseven

func RemoveElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[j] == val {
			if nums[i] == val {
				continue
			} else {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
				j++
			}
		} else {
			j++
		}
	}
	return j
}
