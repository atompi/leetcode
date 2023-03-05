package twofivesevennine

func ColoredCells(n int) int64 {
	return int64((2*n - 1) + 2*(n-1)*(n-1))
}
