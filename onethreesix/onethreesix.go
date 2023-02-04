package onethreesix

func SingleNumber(nums []int) (ans int) {
	for _, v := range nums {
		ans ^= v
	}
	return
}
