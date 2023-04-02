package lc589

type Node struct {
	Val      int
	Children []*Node
}

func Preorder(root *Node) []int {
	res := []int{}
	preorderRecursive(root, &res)
	return res
}

func preorderRecursive(root *Node, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		for i := 0; i < len(root.Children); i++ {
			preorderRecursive(root.Children[i], res)
		}
	}
}
