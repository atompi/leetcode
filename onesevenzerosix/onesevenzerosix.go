package onesevenzerosix

func FindBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	ans := []int{}
	for col := 0; col < n; col++ {
		curCol := col
		for curRow := 0; curRow < m; curRow++ {
			nextCol := curCol + grid[curRow][curCol]
			if nextCol < 0 || nextCol >= n || grid[curRow][curCol] != grid[curRow][nextCol] {
				curCol = -1
				break
			}
			curCol = nextCol
		}
		ans = append(ans, curCol)
	}
	return ans
}
