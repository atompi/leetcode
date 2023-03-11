package lc202

func IsHappy(n int) bool {
	sumMap := make(map[int]int)
	for {
		posArr := make([]int, 0)
		for n > 0 {
			posArr = append(posArr, n%10)
			n /= 10
		}
		sum := 0
		for i := 0; i < len(posArr); i++ {
			sum += posArr[i] * posArr[i]
		}
		if sum == 1 {
			return true
		}
		if _, ok := sumMap[sum]; ok {
			return false
		} else {
			sumMap[sum] = 1
			n = sum
		}
	}
}
