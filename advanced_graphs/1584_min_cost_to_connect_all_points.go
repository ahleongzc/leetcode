package advancedgraphs

import (
	"container/heap"
	"math"
)

type minHeap2 []edge2

func (h minHeap2) Len() int           { return len(h) }
func (h minHeap2) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h minHeap2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap2) Push(x any) {
	*h = append(*h, x.(edge2))
}

func (h *minHeap2) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type edge2 struct {
	weight, node, neighbour int
}

func MinCostConnectPoints(points [][]int) int {
	cost := 0
	distance := func(x1, x2, y1, y2 int) int {
		return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
	}

	pq := minHeap2{}
	for i, dst := range points[1:] {
		edge := edge2{
			weight:    distance(points[0][0], dst[0], points[0][1], dst[1]),
			node:      0,
			neighbour: i + 1,
		}
		heap.Push(&pq, edge)
	}

	visited := make(map[int]struct{})
	visited[0] = struct{}{}

	for len(visited) != len(points) {
		element := heap.Pop(&pq).(edge2)
		next := element.neighbour
		if _, ok := visited[next]; ok {
			continue
		}

		visited[next] = struct{}{}
		cost += element.weight

		for i, point := range points {
			if i == next {
				continue
			}
			if _, ok := visited[i]; ok {
				continue
			}

			edge := edge2{
				weight:    distance(points[next][0], point[0], points[next][1], point[1]),
				node:      next,
				neighbour: i,
			}
			heap.Push(&pq, edge)
		}
	}
	return cost
}
