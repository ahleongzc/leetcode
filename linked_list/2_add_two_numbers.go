package linkedlist

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	ptrToHead := head
	carryOver := 0

	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carryOver
		insertVal := sum % 10
		carryOver = sum / 10

		head.Next = &ListNode{
			Val: insertVal,
		}
		head = head.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + carryOver
		insertVal := sum % 10
		carryOver = sum / 10

		head.Next = &ListNode{
			Val: insertVal,
		}
		head = head.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + carryOver
		insertVal := sum % 10
		carryOver = sum / 10

		head.Next = &ListNode{
			Val: insertVal,
		}
		head = head.Next
		l2 = l2.Next
	}

	if carryOver != 0 {
		head.Next = &ListNode{
			Val: carryOver,
		}
	}

	return ptrToHead.Next
}
