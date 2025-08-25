package arraysandhashing

import (
	"container/heap"
)

type customStruct struct {
	count   int
	element int
}

type pq []customStruct

func (h pq) Len() int           { return len(h) }
func (h pq) Less(i, j int) bool { return h[i].count > h[j].count }
func (h pq) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *pq) Push(x any) {
	*h = append(*h, x.(customStruct))
}

func (h *pq) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TopKFrequentPQ(nums []int, k int) []int {
	pq := &pq{}
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
	}

	for k, v := range countMap {
		cs := customStruct{
			element: k,
			count:   v,
		}
		heap.Push(pq, cs)
	}

	res := make([]int, k)

	for i := range k {
		cs := heap.Pop(pq).(customStruct)
		res[i] = cs.element
	}

	return res
}

func TopKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	countArr := make([][]int, len(nums)+1)
	for k, v := range countMap {
		if countArr[v] == nil {
			countArr[v] = make([]int, 0)
		}
		countArr[v] = append(countArr[v], k)
	}

	res := make([]int, 0)

	for i := len(nums); i > 0; i-- {
		if len(res) == k {
			break
		}

		if countArr[i] == nil {
			continue
		}

		res = append(res, countArr[i]...)
	}

	return res
}
