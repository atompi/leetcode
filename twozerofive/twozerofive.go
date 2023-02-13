package twozerofive

func IsIsomorphic(s string, t string) bool {
	stringByte := []byte(t)
	patternByte := []byte(s)
	stringMap := map[byte]byte{}
	patternMap := map[byte]byte{}

	for index, b := range patternByte {
		if _, ok := patternMap[b]; !ok {
			if _, ok = stringMap[stringByte[index]]; !ok {
				patternMap[b] = stringByte[index]
				stringMap[stringByte[index]] = b
			} else {
				if stringMap[stringByte[index]] != b {
					return false
				}
			}
		} else {
			if patternMap[b] != stringByte[index] {
				return false
			}
		}
	}

	return true
}
