package linkedlist

func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	findMiddle := func(head *ListNode) *ListNode {
		slow := head
		fast := head.Next

		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		}

		return slow
	}

	middle := findMiddle(head)
	secondHalf := middle.Next
	middle.Next = nil

	reverse := func(head *ListNode) *ListNode {
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

	secondHalf = reverse(secondHalf)

	ptr1 := head
	ptr2 := secondHalf

	for ptr1 != nil && ptr2 != nil {
		tmp1 := ptr1.Next
		tmp2 := ptr2.Next

		ptr1.Next = ptr2
		ptr2.Next = tmp1

		ptr1 = tmp1
		ptr2 = tmp2
	}
}
