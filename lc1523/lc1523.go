package lc1523

func CountOdds(low int, high int) int {
	result := 0
	if low == high && low%2 == 1 {
		return 1
	}
	if low%2 == 0 && high%2 == 0 {
		result = (high - low) / 2
	} else if low%2 == 1 && high%2 == 1 {
		result = (high-low-1)/2 + 2
	} else {
		result = (high-low-1)/2 + 1
	}
	return result
}
