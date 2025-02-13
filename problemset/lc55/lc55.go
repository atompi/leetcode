package lc55

func CanJump(nums []int) bool {
	mx := 0
	for i, v := range nums {
		if mx < i {
			return false
		}
		mx = max(mx, i+v)
	}
	return true
}
