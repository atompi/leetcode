package lc2578

import (
	"sort"
)

func SplitNum(num int) int {
	numArr := []int{}
	for num > 0 {
		item := num % 10
		numArr = append(numArr, item)
		num /= 10
	}
	sort.Ints(numArr)
	if len(numArr)%2 != 0 {
		numArr = append([]int{0}, numArr...)
	}
	ans := 0
	for i, j := 0, 0; i < len(numArr); i, j = i+2, j+1 {
		num1 := numArr[i] * power(10, len(numArr)/2-j-1)
		num2 := numArr[i+1] * power(10, len(numArr)/2-j-1)
		ans += num1 + num2
	}
	return ans
}

func power(n, i int) int {
	ans := 1
	for ; i > 0; i-- {
		ans *= n
	}
	return ans
}
