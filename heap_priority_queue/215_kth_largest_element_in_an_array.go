package heappriorityqueue

import "container/heap"

type findKthLargestHeap []int

func (f findKthLargestHeap) Len() int           { return len(f) }
func (f findKthLargestHeap) Less(i, j int) bool { return f[i] < f[j] }
func (f findKthLargestHeap) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f *findKthLargestHeap) Push(x any) {
	*f = append(*f, x.(int))
}

func (f *findKthLargestHeap) Pop() any {
	old := *f
	n := len(old)
	x := old[n-1]
	*f = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	pq := &findKthLargestHeap{}
	*pq = nums
	for i := range *pq {
		(*pq)[i] *= -1
	}

	heap.Init(pq)

	for range k - 1 {
		heap.Pop(pq)
	}

	res := heap.Pop(pq)
	return -1 * res.(int)
}
