package lc2582

func PassThePillow(n int, time int) int {
	// Calculate the number of complete cycles and the remaining time
	cycleLen := 2 * (n - 1)
	// numCycles := time / cycleLen
	remainingTime := time % cycleLen

	// Calculate the index of the person holding the pillow
	if remainingTime < n-1 {
		return remainingTime + 1
	} else {
		return 2*n - remainingTime - 1
	}
}
