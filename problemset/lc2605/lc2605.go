package lc2605

import "sort"

func MinNumber(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i := 0
	j := 0
	allMap := make(map[int]int, len(nums1)+len(nums2))
	for ; i < len(nums1); i++ {
		allMap[nums1[i]] = 1
	}
	for ; j < len(nums2); j++ {
		if allMap[nums2[j]] == 0 {
			allMap[nums2[j]] = 1
		} else {
			return nums2[j]
		}
	}
	a, b := nums1[0], nums2[0]
	swap(&a, &b)
	return a*10 + b
}

func swap(a, b *int) {
	if *a > *b {
		*a, *b = *b, *a
	}
}
