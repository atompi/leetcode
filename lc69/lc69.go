package lc69

// 二分查找
func MySqrt(x int) int {
	low, high := 0, x

	for low <= high {
		mid := (low + high) / 2

		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}

// 牛顿法
func MySqrt2(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}
