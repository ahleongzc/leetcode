package linkedlist

import "math"

func mergeKLists(lists []*ListNode) *ListNode {
	pointers := make([]*ListNode, len(lists))
	copy(pointers, lists)

	res := &ListNode{}
	movingPtr := res

	for {
		currSmallestIndex := -1
		currSmallestVal := math.MaxInt

		for i, pointer := range pointers {
			// linked list processed
			if pointer == nil {
				continue
			}
			if pointer.Val < currSmallestVal {
				currSmallestIndex = i
				currSmallestVal = pointer.Val
			}
		}

		if currSmallestIndex == -1 {
			break
		}

		newListNode := &ListNode{
			Val:  currSmallestVal,
			Next: nil,
		}

		movingPtr.Next = newListNode
		movingPtr = movingPtr.Next
		pointers[currSmallestIndex] = pointers[currSmallestIndex].Next
	}

	return res.Next
}
