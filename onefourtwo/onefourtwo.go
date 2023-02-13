package onefourtwo

type ListNode struct {
	Val  int
	Next *ListNode
}

func DetectCycle(head *ListNode) *ListNode {
	fastPoint := head
	slowPoint := head

	for fastPoint != nil && fastPoint.Next != nil {
		slowPoint = slowPoint.Next
		fastPoint = fastPoint.Next.Next
		if fastPoint == slowPoint {
			pos1 := fastPoint
			pos2 := head
			for pos1 != pos2 {
				pos1 = pos1.Next
				pos2 = pos2.Next
			}
			return pos2
		}
	}
	return nil
}
