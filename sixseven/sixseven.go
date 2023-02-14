package sixseven

import (
	"strconv"
)

func AddBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	result := ""

	for {
		if j >= 0 && i >= 0 {
			intA, _ := strconv.Atoi(string(a[i]))
			intB, _ := strconv.Atoi(string(b[j]))
			c := intA + intB + carry
			carry = c / 2
			result = strconv.Itoa(c%2) + result
			i--
			j--
		} else if i < 0 && j >= 0 {
			intB, _ := strconv.Atoi(string(b[j]))
			c := intB + carry
			carry = c / 2
			result = strconv.Itoa(c%2) + result
			j--
		} else if i >= 0 && j < 0 {
			intA, _ := strconv.Atoi(string(a[i]))
			c := intA + carry
			carry = c / 2
			result = strconv.Itoa(c%2) + result
			i--
		} else {
			if carry == 1 {
				result = strconv.Itoa(carry) + result
			}
			break
		}
	}

	return result
}
