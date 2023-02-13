package nineeight

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nodesList []int

func IsValidBST(root *TreeNode) bool {
	nodesList = []int{} // 初始化 nodesList 很重要
	traversal(root)
	for i := 1; i < len(nodesList); i++ {
		if nodesList[i] <= nodesList[i-1] {
			return false
		}
	}
	return true
}

func traversal(root *TreeNode) {
	if root == nil {
		return
	}
	traversal(root.Left)
	nodesList = append(nodesList, root.Val)
	traversal(root.Right)
}
