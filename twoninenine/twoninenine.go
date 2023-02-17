package twoninenine

import (
	"strconv"
)

func GetHint(secret string, guess string) string {
	a, b := 0, 0
	secretByte := []byte(secret)
	guessByte := []byte(guess)
	secretMap := make(map[byte]int)
	for i := 0; i < len(secretByte); i++ {
		if secretByte[i] == guessByte[i] {
			a++
			guessByte[i] = byte('x')
		} else {
			secretMap[secretByte[i]] += 1
		}
	}
	for i := 0; i < len(guessByte); i++ {
		if _, ok := secretMap[guessByte[i]]; ok {
			if secretMap[guessByte[i]] > 0 {
				b++
				secretMap[guessByte[i]] -= 1
			} else {
				continue
			}
		}
	}
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}
