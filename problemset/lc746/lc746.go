package lc746

func MinCostClimbingStairs(cost []int) int {
	dp0 := cost[0]
	dp1 := cost[1]
	for i := 2; i < len(cost); i++ {
		dpi := min(dp0, dp1) + cost[i]
		dp0, dp1 = dp1, dpi
	}
	return min(dp0, dp1)
}
