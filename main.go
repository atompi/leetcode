package main

import (
	"fmt"

	"github.com/atompi/leetcode/problemset/lc881"
)

func main() {
	people := []int{5, 1, 2, 4}
	limit := 6
	result := lc881.NumRescueBoats(people, limit)
	fmt.Println(result)
}
