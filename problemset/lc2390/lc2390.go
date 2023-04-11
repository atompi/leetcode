package lc2390

func RemoveStars(s string) string {
	sStack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			sStack = sStack[:len(sStack)-1]
		} else {
			sStack = append(sStack, s[i])
		}
	}
	return string(sStack)
}
