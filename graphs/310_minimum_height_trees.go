package graphs

func FindMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	adjList := make(map[int][]int)
	for i := range n {
		adjList[i] = make([]int, 0)
	}
	for _, edge := range edges {
		nodeA, nodeB := edge[0], edge[1]
		adjList[nodeA] = append(adjList[nodeA], nodeB)
		adjList[nodeB] = append(adjList[nodeB], nodeA)
	}

	edgesCountMap := make(map[int]int)
	leaves := make([]int, 0)
	for node, neighbours := range adjList {
		edgesCountMap[node] = len(neighbours)
		if len(neighbours) == 1 {
			leaves = append(leaves, node)
		}
	}

	for len(leaves) != 0 {
		if n <= 2 {
			res := make([]int, 0)
			res = append(res, leaves...)
			return res
		}

		lengthOfLeaves := len(leaves)
		for range lengthOfLeaves {
			leaf := leaves[0]
			leaves = leaves[1:]
			n--

			for _, neighbour := range adjList[leaf] {
				edgesCountMap[neighbour]--

				if edgesCountMap[neighbour] == 1 {
					leaves = append(leaves, neighbour)
				}
			}
		}
	}

	return nil
}
