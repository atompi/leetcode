package main

import (
	"fmt"

	"github.com/atompi/leetcode/threethree"
)

func main() {
	// 2568 out of memory
	// nums := []int{8388608, 131072, 128, 2097152, 65536, 2048, 438, 1048576, 8192, 32, 8, 64, 1024, 2244, 512, 262144, 4096, 16384, 4, 256, 2, 4194304, 2203, 16, 32768, 410, 524288, 765, 1}
	// result := twofivesixeight.MinUnex(nums)

	// 2569 wrong answer
	// nums1 := []int{1, 0, 1}
	// nums2 := []int{0, 0, 0}
	// queries := [][]int{{1, 1, 1}, {2, 1, 0}, {3, 0, 0}}
	// result := twofivesixnine.AnswerQueries(nums1, nums2, queries)

	// 2567 wrong answer
	// nums := []int{1, 4, 3}
	// result := twofivesixseven.MinimumScore(nums)

	nums := []int{1, 2, 3, 0, 21, 22}
	target := 21
	result := threethree.Search(nums, target)
	fmt.Println(result)
}
