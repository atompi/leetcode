package lc45

func Jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	choose, reach, steps := 0, 0, 0
	for i, x := range nums {
		if i+x > reach {
			reach = i + x
			if reach >= len(nums)-1 {
				return steps + 1
			}
		}
		if i == choose {
			choose = reach
			steps++
		}
	}
	return steps
}
