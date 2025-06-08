package linkedlist

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	prev := dummy
	for range left - 1 {
		prev = prev.Next
	}

	curr := prev.Next
	reversePrev := &ListNode{
		Next: curr,
	}
	for range right - left + 1 {
		next := curr.Next
		curr.Next = reversePrev
		reversePrev = curr
		curr = next
	}
	prev.Next.Next = curr
	prev.Next = reversePrev

	return dummy.Next
}
