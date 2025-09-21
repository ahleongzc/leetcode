package twoddp

import (
	"container/heap"
	"math"
)

type customStruct struct {
	diff int
	row  int
	col  int
}

type customStructHeap []customStruct

func (c customStructHeap) Len() int           { return len(c) }
func (c customStructHeap) Less(i, j int) bool { return c[i].diff < c[j].diff }
func (c customStructHeap) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c *customStructHeap) Push(x any) {
	*c = append(*c, x.(customStruct))
}

func (c *customStructHeap) Pop() any {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[0 : n-1]
	return x
}

func minimumEffortPath(heights [][]int) int {
	LENGTH, HEIGHT := len(heights), len(heights[0])
	pq := &customStructHeap{{0, 0, 0}}
	visited := make(map[[2]int]struct{})
	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for pq.Len() > 0 {
		cs := heap.Pop(pq).(customStruct)
		diff, row, col := cs.diff, cs.row, cs.col

		if _, ok := visited[[2]int{row, col}]; ok {
			continue
		}
		visited[[2]int{row, col}] = struct{}{}

		if row == LENGTH-1 && col == HEIGHT-1 {
			return diff
		}

		for _, element := range directions {
			dx, dy := element[0], element[1]
			newRow, newCol := row+dx, col+dy

			if newRow == LENGTH || newRow < 0 || newCol == HEIGHT || newCol < 0 {
				continue
			}

			if _, ok := visited[[2]int{newRow, newCol}]; ok {
				continue
			}

			newDiff := max(
				diff, int(
					math.Abs(
						float64(heights[row][col]-heights[newRow][newCol]),
					),
				),
			)

			heap.Push(pq, customStruct{
				newDiff,
				newRow,
				newCol,
			})
		}
	}

	return 0
}
