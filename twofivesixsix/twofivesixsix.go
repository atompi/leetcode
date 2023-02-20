package twofivesixsix

import (
	"strconv"
)

func MinMaxDifference(num int) int {
	numString := strconv.Itoa(num)
	numArr := []rune(numString)
	firstMinNum := numArr[0]
	firstMaxNum := numArr[0]
	for _, v := range numArr {
		if v != '9' {
			firstMaxNum = v
			break
		}
	}
	maxChart := '9'
	minChart := '0'
	maxNumArr, minNumArr := make([]rune, len(numArr)), make([]rune, len(numArr))
	for i := 0; i < len(numArr); i++ {
		if numArr[i] == firstMinNum {
			minNumArr[i] = minChart
		} else {
			minNumArr[i] = numArr[i]
		}
	}
	for i := 0; i < len(numArr); i++ {
		if numArr[i] == firstMaxNum {
			maxNumArr[i] = maxChart
		} else {
			maxNumArr[i] = numArr[i]
		}
	}
	maxNum, _ := strconv.Atoi(string(maxNumArr))
	minNum, _ := strconv.Atoi(string(minNumArr))
	return maxNum - minNum
}
