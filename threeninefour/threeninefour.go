package threeninefour

import (
	"strings"
	"unicode"
)

func DecodeString(s string) string {
	stack := make([]interface{}, 0)
	currentStr := ""
	currentNum := 0
	for _, c := range s {
		if unicode.IsDigit(c) {
			currentNum = currentNum*10 + int(c-'0')
		} else if c == '[' {
			stack = append(stack, currentStr)
			stack = append(stack, currentNum)
			currentStr = ""
			currentNum = 0
		} else if c == ']' {
			num := stack[len(stack)-1].(int)
			stack = stack[:len(stack)-1]
			prevStr := stack[len(stack)-1].(string)
			stack = stack[:len(stack)-1]
			currentStr = prevStr + strings.Repeat(currentStr, num)
		} else {
			currentStr += string(c)
		}
	}
	return currentStr
}
