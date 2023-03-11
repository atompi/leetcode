package lc1345

// ChatGPT
func MinJumps(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 0
	}

	// Create a map to store the indices of each value in the array
	indexMap := make(map[int][]int)
	for i := 0; i < n; i++ {
		indexMap[arr[i]] = append(indexMap[arr[i]], i)
	}

	// Initialize the queue with the first index of the array
	queue := []int{0}
	visited := make([]bool, n)
	visited[0] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == n-1 {
				return steps
			}

			// Add the adjacent indices to the queue
			if curr+1 < n && !visited[curr+1] {
				queue = append(queue, curr+1)
				visited[curr+1] = true
			}

			if curr-1 >= 0 && !visited[curr-1] {
				queue = append(queue, curr-1)
				visited[curr-1] = true
			}

			// Add the indices with the same value to the queue
			for _, j := range indexMap[arr[curr]] {
				if j != curr && !visited[j] {
					queue = append(queue, j)
					visited[j] = true
				}
			}

			// Clear the indices with the same value from the map
			delete(indexMap, arr[curr])
		}
		steps++
	}

	return -1
}
