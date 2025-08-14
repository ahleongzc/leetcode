package linkedlist

// A -> B -> C -> D -> E

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	slow := dummy
	fast := dummy

	for range n + 1 {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
