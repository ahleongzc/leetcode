package linkedlist

// Assuming the gap is length n
// Every iteration the tortoise will increase the gap by 1, and the hare will decrease the length by 2
// Therefore n - 2 + 1 = n - 1

func HasCycleFloydCycleDetection(head *ListNode) bool {
	tortoise := head
	hare := head

	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next

		if tortoise == hare {
			return true
		}
	}

	return false
}

func HasCycle(head *ListNode) bool {
	history := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := history[head]; ok {
			return true
		}

		history[head] = true
		head = head.Next
	}

	return false
}
