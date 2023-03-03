package twoeight

// 遍历
func StrStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	for i := 0; i <= n-m; i++ {
		if haystack[i:i+m] == needle {
			return i
		}
	}
	return -1
}

// Rabin-Karp 算法
func StrStr2(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	sha, target, left, right, mod := 0, 0, 0, 0, 1<<31-1
	multi := 1
	for i := 0; i < m; i++ {
		target = (target*256%mod + int(needle[i])) % mod
	}
	for i := 1; i < m; i++ {
		multi = multi * 256 % mod
	}

	for ; right < n; right++ {
		sha = (sha*256%mod + int(haystack[right])) % mod
		if right-left+1 < m {
			continue
		}
		// 此时 left~right 的长度已经为 needle 的长度 m 了，只需要比对 sha 值与 target 是否一致即可
		// 为避免 hash 冲突，还需要确保 haystack[left:right+1] 与 needle 相同
		if sha == target && haystack[left:right+1] == needle {
			return left
		}
		// 未匹配成功，left 右移一位
		sha = (sha - (int(haystack[left])*multi)%mod + mod) % mod
		left++
	}
	return -1
}

// KMP
func kmp(s1, s2 string) int {
	if len(s1) < len(s2) {
		return -1
	}
	next := getNext(s2)
	for i, j := 0, 0; i < len(s1); {
		if j == -1 || s1[i] == s2[j] {
			i++
			j++
		} else {
			j = next[j]
		}
		if j == len(s2) {
			return i - j
		}
	}
	return -1
}

func getNext(s string) []int {
	next := make([]int, len(s))
	next[0] = -1
	for i, j := 0, -1; i < len(s)-1; {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			if s[i] != s[j] {
				next[i] = j
			} else {
				next[i] = next[j]
			}
		} else {
			j = next[j]
		}
	}
	return next
}

func getNext2(s string) []int {
	next := make([]int, len(s)+1)
	next[0] = -1
	for i, j := 0, -1; i <= len(s)-1; {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}
