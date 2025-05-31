package main

import (
	"fmt"
	graphs "leetcode/graphs"
)

func main() {
	edges := []int{2, 2, 3, -1}
	fmt.Println(graphs.ClosestMeetingNode(edges, 0, 1))

	edges = []int{1, 2, -1}
	fmt.Println(graphs.ClosestMeetingNode(edges, 0, 2))
}
