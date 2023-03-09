package lc58

import (
	"strings"
)

// 双指针找出所有单词的长度
func LengthOfLastWord(s string) int {
	j := 0
	sByte := strings.Split(s, "")
	sByte = append(sByte, " ")
	wordLength := []int{}
	for i := 0; i < len(sByte); i++ {
		if (sByte[i] == " " && sByte[j] == " ") || (sByte[i] != " " && sByte[j] != " ") {
			continue
		} else if sByte[i] == " " && sByte[j] != " " {
			curLength := i - j
			j = i
			wordLength = append(wordLength, curLength)
		} else if sByte[i] != " " && sByte[j] == " " {
			j = i
		}
	}
	return wordLength[len(wordLength)-1]
}

// 只处理最后一个单词
func LengthOfLastWord2(s string) int {
	i := len(s) - 1
	res := 0

	for i >= 0 {
		if s[i] == ' ' {
			i--
			continue
		}
		break
	}

	for i >= 0 {
		if s[i] == ' ' {
			break
		}
		i--
		res++
	}

	return res
}

// strings.Fields 方法切割字符串
func LengthOfLastWord3(s string) int {
	str := strings.Fields(s)
	return len(str[len(str)-1])
}
