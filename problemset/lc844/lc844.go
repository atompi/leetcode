package lc844

func BackspaceCompare(s string, t string) bool {
	bcS := backspaceString(s)
	bcT := backspaceString(t)
	return bcS == bcT
}

func backspaceString(s string) string {
	rawStack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			rawStack = append(rawStack, s[i])
		} else {
			if len(rawStack)-1 >= 0 {
				rawStack = rawStack[0 : len(rawStack)-1]
			} else {
				continue
			}
		}
	}
	return string(rawStack)
}
