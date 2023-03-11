package lc102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		curQueueSize := len(queue)
		vec := make([]int, 0, curQueueSize)
		for i := 0; i < curQueueSize; i++ {
			vec = append(vec, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[curQueueSize:]
		result = append(result, vec)
	}

	return result
}
