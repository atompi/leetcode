package lc783

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MinDiffInBST(root *TreeNode) int {
	ans, preVal := math.MaxInt16, -1
	traversal(root, &ans, &preVal)
	return ans
}

func traversal(curNode *TreeNode, ans, preVal *int) {
	if curNode == nil {
		return
	}
	traversal(curNode.Left, ans, preVal)
	if *preVal != -1 {
		*ans = min(*ans, abs(curNode.Val-*preVal))
	}
	*preVal = curNode.Val
	traversal(curNode.Right, ans, preVal)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
