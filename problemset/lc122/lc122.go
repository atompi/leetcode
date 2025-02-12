package lc122

func MaxProfit(prices []int) int {
	c := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] > prices[i+1] {
			continue
		} else {
			c += prices[i+1] - prices[i]
		}
	}
	return c
}
