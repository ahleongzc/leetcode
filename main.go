package main

import (
	"fmt"
	graphs "leetcode/graphs"
)

func main() {
	grid := [][]int{{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}}

	fmt.Println(graphs.IslandPerimeter(grid))
}
