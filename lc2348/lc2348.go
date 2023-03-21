package lc2348

func ZeroFilledSubarray(nums []int) int64 {
	z := 0
	var sum int64 = 0
	for _, n := range nums {
		if n == 0 {
			z++
		} else {
			sum += int64(z * (z + 1) / 2)
			z = 0
		}
	}
	return sum + int64(z*(z+1)/2)
}
