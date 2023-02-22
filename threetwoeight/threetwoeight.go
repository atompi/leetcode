package threetwoeight

type ListNode struct {
	Val  int
	Next *ListNode
}

func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var oddHead, evenHead, oddTail, evenTail *ListNode

	isOdd := true

	for head != nil {
		next := head.Next
		head.Next = nil

		if isOdd {
			if oddHead == nil {
				oddHead = head
				oddTail = head
			} else {
				oddTail.Next = head
				oddTail = head
			}
		} else {
			if evenHead == nil {
				evenHead = head
				evenTail = head
			} else {
				evenTail.Next = head
				evenTail = head
			}
		}

		isOdd = !isOdd

		head = next
	}

	oddTail.Next = evenHead

	return oddHead
}
