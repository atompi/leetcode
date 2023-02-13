package onesixtwosix

import (
	"sort"

	"github.com/atompi/leetcode/utils/max"
)

func BestTeamScore(scores []int, ages []int) int {
	n := len(ages)
	player := [][]int{}
	for i := 0; i < n; i++ {
		player = append(player, []int{ages[i], scores[i]})
	}
	sort.Slice(player, func(i, j int) bool {
		return player[i][0] < player[j][0] || (player[i][0] == player[j][0] && player[i][1] < player[j][1])
	})
	result := 0
	DP := make([]int, n+1)
	for i := 0; i < n; i++ {
		DP[i+1] = player[i][1]
		for j := 0; j < i; j++ {
			if player[j][1] > player[i][1] {
				continue
			}
			DP[i+1] = max.MaxOfTwoInt(DP[i+1], player[i][1]+DP[j+1])
		}
		result = max.MaxOfTwoInt(result, DP[i+1])
	}
	return result
}
