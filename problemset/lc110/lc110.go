package lc110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	return getDepth(root) != -1
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftDepth := getDepth(node.Left)
	if leftDepth == -1 {
		return -1
	}
	rightDeph := getDepth(node.Right)
	if rightDeph == -1 {
		return -1
	}
	if abs(leftDepth-rightDeph) > 1 {
		return -1
	} else {
		return 1 + max(leftDepth, rightDeph)
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
