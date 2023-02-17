package main

import (
	"fmt"

	"github.com/atompi/leetcode/twoninenine"
)

func main() {
	secret := "1"
	guess := "1"
	result := twoninenine.GetHint(secret, guess)
	fmt.Println(result)
}
