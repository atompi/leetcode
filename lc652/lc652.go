package lc652

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var ans []*TreeNode
	cnt := make(map[string]int)
	var dfs func(root *TreeNode) string

	dfs = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		v := strconv.Itoa(root.Val) + "," + dfs(root.Left) + "," + dfs(root.Right)
		cnt[v]++
		if cnt[v] == 2 {
			ans = append(ans, root)
		}
		return v
	}

	dfs(root)

	return ans
}
