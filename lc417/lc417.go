package lc417

func PacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacific, atlantic := make([][]bool, m), make([][]bool, m)
	result := make([][]int, 0)

	// Initialize pacific and atlantic to false
	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	// DFS from all border cells
	for i := 0; i < m; i++ {
		dfs(heights, i, 0, pacific)
		dfs(heights, i, n-1, atlantic)
	}
	for j := 0; j < n; j++ {
		dfs(heights, 0, j, pacific)
		dfs(heights, m-1, j, atlantic)
	}

	// Find cells reachable from both oceans
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func dfs(heights [][]int, i, j int, ocean [][]bool) {
	ocean[i][j] = true
	for _, dir := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		x, y := i+dir[0], j+dir[1]
		if x >= 0 && x < len(heights) && y >= 0 && y < len(heights[0]) && !ocean[x][y] && heights[x][y] >= heights[i][j] {
			dfs(heights, x, y, ocean)
		}
	}
}
