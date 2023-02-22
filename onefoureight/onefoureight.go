package onefoureight

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMiddle(head)
	sortedFromMiddle := SortList(mid.Next)
	mid.Next = nil
	sortedTilMiddle := SortList(head)

	return merge(sortedTilMiddle, sortedFromMiddle)
}

func merge(head1, head2 *ListNode) *ListNode {
	tmp := ListNode{math.MinInt32, nil}
	cur := &tmp

	for ; head1 != nil && head2 != nil; cur = cur.Next {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
	}

	if head1 == nil && head2 == nil {
		return tmp.Next
	}

	if head1 != nil {
		cur.Next = head1
	} else {
		cur.Next = head2
	}

	cur = tmp.Next

	return cur
}

func getMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
