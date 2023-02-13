package eightsevensix

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	front := head
	back := head
	for {
		if front.Next != nil && front.Next.Next != nil {
			back = back.Next
			front = front.Next.Next
		} else if front.Next != nil && front.Next.Next == nil {
			back = back.Next
			front = front.Next
		} else {
			break
		}
	}
	return back
}
