package lc43

import (
	"strconv"
)

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	tmpVal := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			tmpVal[i+j+1] += mul
			tmpVal[i+j] += tmpVal[i+j+1] / 10
			tmpVal[i+j+1] = tmpVal[i+j+1] % 10
		}
	}
	ans := ""
	i := 0
	for i < len(tmpVal) && tmpVal[i] == 0 {
		i++
	}
	for ; i < len(tmpVal); i++ {
		ans += strconv.Itoa(tmpVal[i])
	}
	return ans
}
