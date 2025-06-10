package heappriorityqueue

import "container/heap"

type minHeap []int

func (mh minHeap) Len() int {
	return len(mh)
}

func (mh minHeap) Less(i, j int) bool {
	return mh[i] < mh[j]
}

func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minHeap) Push(x any) {
	element, ok := x.(int)
	if !ok {
		return
	}
	*mh = append(*mh, element)
}

func (mh *minHeap) Pop() any {
	element := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return element
}

func FindKthLargest(nums []int, k int) int {
	numsReversed := make([]int, len(nums))
	for i, num := range nums {
		numsReversed[i] = -1 * num
	}

	mh := &minHeap{}
	*mh = numsReversed
	heap.Init(mh)

	for range k - 1 {
		heap.Pop(mh)
	}

	return heap.Pop(mh).(int) * -1
}
