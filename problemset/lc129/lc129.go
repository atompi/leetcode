package lc129

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumNumbers(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, s int) int {
		if root == nil {
			return 0
		}
		s = s*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return s
		}
		return dfs(root.Left, s) + dfs(root.Right, s)
	}
	return dfs(root, 0)
}
