package lc2300

import "sort"

func SuccessfulPairs(spells []int, potions []int, success int64) []int {
	ans := []int{}
	sort.Ints(potions)
	m := len(potions)
	for _, v := range spells {
		i := sort.Search(m, func(i int) bool { return int64(potions[i]*v) >= success })
		ans = append(ans, m-i)
	}
	return ans
}
