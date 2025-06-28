package advancedgraphs

import (
	"container/heap"
)

type minHeap []edge

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(edge))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type edge struct {
	node, weight int
}

func NetworkDelayTime(times [][]int, n int, k int) int {
	adjList := make(map[int][]edge)
	for _, time := range times {
		source := time[0]
		target := edge{node: time[1], weight: time[2]}
		adjList[source] = append(adjList[source], target)
	}

	pq := &minHeap{}
	heap.Push(pq, edge{
		node:   k,
		weight: 0,
	})

	visited := make(map[int]bool)
	totalTime := 0

	for pq.Len() != 0 {
		item := heap.Pop(pq).(edge)
		node, time := item.node, item.weight

		if visited[node] {
			continue
		}
		visited[node] = true
		totalTime = time

		for _, neighbour := range adjList[node] {
			if visited[neighbour.node] {
				continue
			}
			heap.Push(pq, edge{
				node:   neighbour.node,
				weight: neighbour.weight + time,
			})
		}
	}

	if len(visited) == n {
		return totalTime
	}

	return -1
}
