package main

import (
	"fmt"

	"github.com/atompi/leetcode/problemset/lc123"
)

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	result := lc123.MaxProfit(prices)
	fmt.Println(result)
}
