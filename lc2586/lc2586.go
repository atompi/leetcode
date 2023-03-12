package lc2586

func VowelStrings(words []string, left int, right int) int {
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	count := 0
	for i := left; i <= right; i++ {
		for j := 0; j < len(vowels); j++ {
			if words[i][0] == vowels[j] {
				for k := 0; k < len(vowels); k++ {
					if words[i][len(words[i])-1] == vowels[k] {
						count++
						break
					}
				}
				break
			}
		}
	}
	return count
}
