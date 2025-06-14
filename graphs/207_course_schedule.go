package graphs

func CanFinish(numCourses int, prerequisites [][]int) bool {
	// Initialising the adjacency list
	adjList := make(map[int][]int)
	for i := range numCourses {
		adjList[i] = make([]int, 0)
	}

	for _, prerequisitePair := range prerequisites {
		curr, pre := prerequisitePair[0], prerequisitePair[1]
		adjList[curr] = append(adjList[curr], pre)
	}

	visited := make(map[int]bool)
	processing := make(map[int]bool)

	var dfs func(curr int) bool
	dfs = func(curr int) bool {
		if _, exists := visited[curr]; exists {
			return true
		}
		if _, exists := processing[curr]; exists {
			return false
		}

		processing[curr] = true

		for _, neighbour := range adjList[curr] {
			if !dfs(neighbour) {
				return false
			}
		}

		delete(processing, curr)
		visited[curr] = true

		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return false
		}
	}

	return true
}
