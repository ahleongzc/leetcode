package graphs

import (
	"math"
)

func ClosestMeetingNode(edges []int, node1 int, node2 int) int {
	// The key is the node itself, the value is the distance to that node
	distanceMapNode1 := make(map[int]int)
	distanceMapNode1[node1] = 0

	distanceMapNode2 := make(map[int]int)
	distanceMapNode2[node2] = 0

	var dfs = func(node int, edges []int, distanceMap map[int]int) {
		distance := 0

		for edges[node] != -1 {
			node = edges[node]
			distance++

			if _, ok := distanceMap[node]; ok {
				break
			}

			distanceMap[node] = distance
		}
	}

	dfs(node1, edges, distanceMapNode1)
	dfs(node2, edges, distanceMapNode2)

	smallestIndex := -1
	smallestDistance := math.MaxInt

	for i := range len(edges) {
		if _, ok := distanceMapNode1[i]; !ok {
			continue
		}
		if _, ok := distanceMapNode2[i]; !ok {
			continue
		}

		distance := max(distanceMapNode1[i], distanceMapNode2[i])
		if distance < smallestDistance {
			smallestIndex = i
			smallestDistance = distance
		}
	}

	return smallestIndex
}
