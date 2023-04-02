package lc88

import "sort"

// one line
func Merge(nums1 []int, m int, nums2 []int, n int) {
	sort.Ints(append(nums1[:m], nums2...))
}

// 三指针
func Merge2(nums1 []int, m int, nums2 []int, n int) {
	var ptr1, ptr2, ptr3 int = m - 1, n - 1, m + n - 1
	for ; ptr1 >= 0 && ptr2 >= 0; ptr3-- {
		if nums2[ptr2] >= nums1[ptr1] {
			nums1[ptr3] = nums2[ptr2]
			ptr2--
		} else {
			nums1[ptr3] = nums1[ptr1]
			ptr1--
		}
	}

	if ptr2 >= 0 {
		copy(nums1[:ptr3+1], nums2[:ptr2+1])
	}
}
