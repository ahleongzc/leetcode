package advancedgraphs

import (
	"container/heap"
	"math"
)

type grid struct {
	x    int
	y    int
	diff int
}

type minEffortPathHeap []grid

func (h minEffortPathHeap) Len() int           { return len(h) }
func (h minEffortPathHeap) Less(i, j int) bool { return h[i].diff < h[j].diff }
func (h minEffortPathHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minEffortPathHeap) Push(x any)        { *h = append(*h, x.(grid)) }
func (h *minEffortPathHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func MinimumEffortPath(heights [][]int) int {
	rows := len(heights)
	cols := len(heights[0])

	pq := &minEffortPathHeap{}

	heap.Push(pq, grid{
		x:    0,
		y:    0,
		diff: 0,
	})

	directions := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	visited := make(map[[2]int]struct{})

	for pq.Len() != 0 {
		element := heap.Pop(pq).(grid)
		if element.x == rows-1 && element.y == cols-1 {
			return element.diff
		}

		visited[[2]int{element.x, element.y}] = struct{}{}

		for _, dir := range directions {
			newX := element.x + dir[0]
			newY := element.y + dir[1]

			if newX < 0 || newY < 0 || newX == rows || newY == cols {
				continue
			}

			if _, ok := visited[[2]int{newX, newY}]; ok {
				continue
			}

			heap.Push(pq, grid{
				x:    newX,
				y:    newY,
				diff: max(element.diff, int(math.Abs(float64(heights[newX][newY]-heights[element.x][element.y])))),
			})
		}
	}

	return -1
}
