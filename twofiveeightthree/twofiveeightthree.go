package twofiveeightthree

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthLargestLevelSum(root *TreeNode, k int) int64 {
	// Traverse the tree and calculate the level sums
	levelSums := make(map[int]int)
	traverse(root, 0, levelSums)

	// Sort the level sums in descending order
	sortedSums := make([]int, 0, len(levelSums))
	for _, sum := range levelSums {
		sortedSums = append(sortedSums, sum)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedSums)))

	// Return the kth largest level sum if it exists
	if k <= len(sortedSums) {
		return int64(sortedSums[k-1])
	} else {
		return int64(-1)
	}
}

func traverse(node *TreeNode, level int, levelSums map[int]int) {
	if node == nil {
		return
	}

	// Add the node value to the current level sum
	levelSums[level] += node.Val

	// Traverse the left and right subtrees
	traverse(node.Left, level+1, levelSums)
	traverse(node.Right, level+1, levelSums)
}
