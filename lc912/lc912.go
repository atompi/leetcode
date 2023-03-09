package lc912

func SortArray(nums []int) []int {
	l := len(nums)
	heapSort(nums, l)
	return nums
}

func heapify(nums []int, l, i int) {
	max := i
	left := i*2 + 1
	right := i*2 + 2
	if left < l && nums[left] > nums[max] {
		max = left
	}
	if right < l && nums[right] > nums[max] {
		max = right
	}
	if max != i {
		nums[max], nums[i] = nums[i], nums[max]
		heapify(nums, l, max)
	}
}

func heapSort(nums []int, l int) {
	for i := l/2 - 1; i >= 0; i-- {
		heapify(nums, l, i)
	}
	for i := l - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
}
