package lc121

func MaxProfit(prices []int) int {
	curPrice, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-curPrice > profit {
			profit = prices[i] - curPrice
		}
		if prices[i] < curPrice {
			curPrice = prices[i]
		}
	}
	return profit
}
