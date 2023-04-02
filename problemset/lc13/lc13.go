package lc13

func RomanToInt(s string) int {
	romanDict := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	if s == "" {
		return 0
	}

	currentNum, nextNum, result := 0, 0, 0
	for i := 0; i < len(s)-1; i++ {
		currentRomanChar := string(s[i])
		nextRomanChar := string(s[i+1])
		currentNum = romanDict[currentRomanChar]
		nextNum = romanDict[nextRomanChar]
		if currentNum < nextNum {
			result = result - currentNum
		} else {
			result = result + currentNum
		}
	}
	return result + romanDict[string(s[len(s)-1])]
}
