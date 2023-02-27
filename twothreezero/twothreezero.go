package twothreezero

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthSmallest(root *TreeNode, k int) int {
	stk := []*TreeNode{}
	for root != nil || len(stk) > 0 {
		if root != nil {
			stk = append(stk, root)
			root = root.Left
		} else {
			root = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			k--
			if k == 0 {
				return root.Val
			}
			root = root.Right
		}
	}
	return 0
}

// Follow up
type BST struct {
	cnt  map[*TreeNode]int
	root *TreeNode
}

func newBST(root *TreeNode) *BST {
	var count func(*TreeNode) int
	cnt := map[*TreeNode]int{}
	count = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		n := 1 + count(root.Left) + count(root.Right)
		cnt[root] = n
		return n
	}
	count(root)
	return &BST{cnt, root}
}

func (bst *BST) kthSmallest(k int) int {
	node := bst.root
	for node != nil {
		v := 0
		if node.Left != nil {
			v = bst.cnt[node.Left]
		}
		if v == k-1 {
			return node.Val
		}
		if v < k-1 {
			k -= (v + 1)
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return 0
}

func KthSmallest2(root *TreeNode, k int) int {
	bst := newBST(root)
	return bst.kthSmallest(k)
}
