package heappriorityqueue

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	if val, ok := x.(int); ok {
		*h = append(*h, val)
	}
}

func (h *MinHeap) Pop() any {
	element := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return element
}
