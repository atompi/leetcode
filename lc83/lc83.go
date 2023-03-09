package lc83

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	curNode := head
	if curNode == nil {
		return head
	}
	for curNode.Next != nil {
		if curNode.Val == curNode.Next.Val {
			if curNode.Next.Next != nil {
				curNode.Next = curNode.Next.Next
			} else {
				curNode.Next = nil
			}
		} else {
			curNode = curNode.Next
		}
	}
	return head
}
