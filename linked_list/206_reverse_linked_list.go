package linkedlist

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	next := head

	for next != nil {
		tmp := next.Next
		next.Next = prev

		prev = next
		next = tmp
	}

	return prev
}
