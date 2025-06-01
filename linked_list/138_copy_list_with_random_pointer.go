package linkedlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)

	// First pass
	curr := head
	for curr != nil {
		nodeMap[curr] = &Node{
			Val: curr.Val,
		}
		curr = curr.Next
	}

	// Second pass
	curr = head
	for curr != nil {
		if curr.Next != nil {
			nodeMap[curr].Next = nodeMap[curr.Next]
		}
		if curr.Random != nil {
			nodeMap[curr].Random = nodeMap[curr.Random]
		}
		curr = curr.Next
	}

	return nodeMap[head]
}
