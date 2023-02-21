package onenine

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	fast := head
	slow := dummy
	for fast != nil {
		if n > 0 {
			fast = fast.Next
			n--
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
