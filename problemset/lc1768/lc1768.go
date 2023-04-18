package lc1768

func MergeAlternately(word1 string, word2 string) string {
	ans := []byte{}
	// 获取两个字符串的最小值
	min1 := min(len(word1), len(word2))
	// 循环 min1 次
	for i := 0; i < min1; i++ {
		ans = append(ans, word1[i], word2[i])
	}
	ans = append(ans, word1[min1:]...)
	ans = append(ans, word2[min1:]...)
	// 将 ans 转为 string 类型
	return string(ans)
}

// 返回两个整数中的最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
