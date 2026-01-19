package graphs

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjList := make(map[int][]int)
	for _, prerequisite := range prerequisites {
		before, after := prerequisite[1], prerequisite[0]
		adjList[before] = append(adjList[before], after)
	}

	processing := make(map[int]struct{})
	visited := make(map[int]struct{})

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if _, ok := processing[course]; ok {
			return false
		}
		
		if _, ok := visited[course]; ok {
			return true
		}

		processing[course] = struct{}{}

		for _, child := range adjList[course] {
			if !dfs(child) {
				return false
			}
		}

		delete(processing, course)
		visited[course] = struct{}{}
		
		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return false
		}
	}

	return true
}
