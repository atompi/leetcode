package lc994

func OrangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cnt := 0
	var q [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			} else if grid[i][j] == 1 {
				cnt++
			}
		}
	}
	ans := 0
	dirs := []int{-1, 0, 1, 0, -1}
	for len(q) > 0 && cnt > 0 {
		ans++
		for i := len(q); i > 0; i-- {
			p := q[0]
			q = q[1:]
			for j := 0; j < 4; j++ {
				x, y := p[0]+dirs[j], p[1]+dirs[j+1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					cnt--
					grid[x][y] = 2
					q = append(q, []int{x, y})
				}
			}
		}
	}
	if cnt > 0 {
		ans = -1
	}
	return ans
}

// ChatGPT
func OrangesRotting2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := make([][2]int, 0)
	freshCount := 0

	// Add all rotten oranges to the queue
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	// Use BFS to infect adjacent fresh oranges
	minutes := 0
	for len(queue) > 0 {
		for size := len(queue); size > 0; size-- {
			i, j := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, dir := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				x, y := i+dir[0], j+dir[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					freshCount--
					queue = append(queue, [2]int{x, y})
				}
			}
		}
		minutes++
	}

	if freshCount > 0 {
		// There are fresh oranges left that are unreachable
		return -1
	} else {
		// All oranges have been infected
		return minutes - 1
	}
}
