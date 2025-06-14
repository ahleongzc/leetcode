package graphs

func FindOrder(numCourses int, prerequisites [][]int) []int {
	adjList := make(map[int][]int)
	for i := range numCourses {
		adjList[i] = make([]int, 0)
	}
	for _, prerequisite := range prerequisites {
		curr, pre := prerequisite[0], prerequisite[1]
		adjList[curr] = append(adjList[curr], pre)
	}

	visited := make(map[int]struct{})
	processing := make(map[int]struct{})
	processed := make([]int, 0)

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if _, exists := processing[i]; exists {
			return false
		}
		if _, exists := visited[i]; exists {
			return true
		}

		processing[i] = struct{}{}

		for _, neighbour := range adjList[i] {
			if !dfs(neighbour) {
				return false
			}
		}

		delete(processing, i)
		visited[i] = struct{}{}
		processed = append(processed, i)

		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return make([]int, 0)
		}
	}

	return processed
}
