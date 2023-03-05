package twofiveeightfour

import "math"

func FindValidSplit(nums []int) int {
	cnt := make([]map[int]int, len(nums))

	for i, x := range nums {
		cnt[i] = fact(x)
	}

	x := map[int]bool{}
	for i := 0; i < len(nums)-1; i++ {
		for prime := range cnt[i] {
			x[prime] = true
		}

		flag := false
		for j := i + 1; j < len(nums) && !flag; j++ {
			for prime := range cnt[j] {
				if _, ok := x[prime]; ok {
					flag = true
				}
			}
		}

		if !flag {
			return i
		}
	}

	return -1
}

func fact(x int) map[int]int {
	res := make(map[int]int)

	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		for x%i == 0 {
			res[i]++
			x /= i
		}
	}

	if x != 1 {
		res[x]++
	}

	return res
}
