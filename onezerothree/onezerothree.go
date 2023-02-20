package onezerothree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	queue := []*TreeNode{root}
	leftToRight := true
	for len(queue) > 0 {
		levelSize := len(queue)
		levelVals := []int{}
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if leftToRight {
				levelVals = append(levelVals, node.Val)
			} else {
				levelVals = append([]int{node.Val}, levelVals...)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, levelVals)
		leftToRight = !leftToRight
	}
	return result
}
