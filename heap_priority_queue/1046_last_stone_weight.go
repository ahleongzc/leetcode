package heappriorityqueue

import "container/heap"

func LastStoneWeight(stones []int) int {
	stonesReversed := make([]int, 0)
	for _, stone := range stones {
		stonesReversed = append(stonesReversed, -1*stone)
	}

	minHeap := &MinHeap{}
	*minHeap = stonesReversed
	heap.Init(minHeap)

	for len(*minHeap) > 1 {
		largest := heap.Pop(minHeap).(int) * -1
		secondLargest := heap.Pop(minHeap).(int) * -1
		diff := largest - secondLargest

		if diff > 0 {
			heap.Push(minHeap, -1*diff)
		}
	}

	if len(*minHeap) == 0 {
		return 0
	}

	return (*minHeap)[0] * -1
}
