package lc136

func SingleNumber(nums []int) (ans int) {
	for _, v := range nums {
		ans ^= v
	}
	return
}
