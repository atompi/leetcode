package lc2609

func FindTheLongestBalancedSubstring(s string) int {
	ans := 0
	zCnt, oCnt := 0, 0
	n := len(s)
	i := 0

	for i < n {
		if s[i] == '0' {
			zCnt = 0
			for i < n && s[i] == '0' {
				zCnt += 1
				i++
			}
			oCnt = 0
			for i < n && s[i] == '1' {
				oCnt += 1
				i++
			}
			ans = max(min(zCnt, oCnt)*2, ans)
		} else {
			i++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
