package twozerosix

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var tmp, pre *ListNode
	cur := head
	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
