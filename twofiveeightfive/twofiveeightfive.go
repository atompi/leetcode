package twofiveeightfive

const mod = 1e9 + 7

func WaysToReachTarget(target int, types [][]int) int {
	// dp[i][j] represents the number of ways to get j marks using first i types of questions
	dp := make([][]int, len(types)+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
		dp[i][0] = 1
	}
	for i := 1; i <= len(types); i++ {
		count, marks := types[i-1][0], types[i-1][1]
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			for k := 1; k <= count; k++ {
				if j-k*marks >= 0 {
					dp[i][j] += dp[i-1][j-k*marks]
					dp[i][j] %= mod
				}
			}
		}
	}
	return dp[len(types)][target]
}
