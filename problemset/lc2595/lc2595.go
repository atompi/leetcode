package lc2595

func EvenOddBit(n int) []int {
	binary := []int{}
	even, odd := 0, 0
	for n > 0 {
		item := n % 2
		n /= 2
		binary = append(binary, item)
	}
	for i := 0; i < len(binary); i++ {
		if i%2 == 0 {
			if binary[i] == 1 {
				even++
			}
		} else {
			if binary[i] == 1 {
				odd++
			}
		}
	}
	return []int{even, odd}
}
