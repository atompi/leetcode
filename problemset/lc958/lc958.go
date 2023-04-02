package lc958

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	for q[0] != nil {
		root = q[0]
		q = q[1:]
		q = append(q, root.Left)
		q = append(q, root.Right)
	}
	for len(q) > 0 && q[0] == nil {
		q = q[1:]
	}
	return len(q) == 0
}
