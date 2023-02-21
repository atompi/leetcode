package twothreefour

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrome(head *ListNode) bool {
	middle := twoPointFindMiddle(head)
	rightListHead := reverse(middle.Next)
	ans := true
	pos1 := head
	pos2 := rightListHead
	for ans && (pos2 != nil) {
		if pos1.Val != pos2.Val {
			ans = false
		}
		pos1 = pos1.Next
		pos2 = pos2.Next
	}
	middle.Next = reverse(rightListHead)
	return ans
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	var next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func twoPointFindMiddle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
