package graphs

func countComponentsDFS(n int, edges [][]int) int {
	visited := make(map[int]struct{})
	adjList := make(map[int][]int)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	var dfs func(node int)
	dfs = func(node int) {
		if _, ok := visited[node]; ok {
			return
		}

		visited[node] = struct{}{}
		for _, children := range adjList[node] {
			dfs(children)
		}
	}

	output := 0
	for i := range n {
		if _, ok := visited[i]; !ok {
			output++
			dfs(i)
		}
	}

	return output
}

func CountComponents(n int, edges [][]int) int {
	parent := make(map[int]int)
	rank := make(map[int]int)

	for i := range n {
		parent[i] = i
		rank[i] = 0
	}

	find := func(i int) int {
		p := parent[i]
		for p != parent[p] {
			parent[p] = parent[parent[p]]
			p = parent[p]
		}
		return p
	}

	union := func(a, b int) {
		parentA, parentB := find(a), find(b)
		if parentA == parentB {
			return
		}

		if rank[parentA] > rank[parentB] {
			parent[parentB] = parentA
		} else {
			parent[parentA] = parentB
			rank[parentA]++
		}
	}

	for _, edge := range edges {
		union(edge[0], edge[1])
		union(edge[1], edge[0])
	}

	uniqueSet := make(map[int]struct{})

	for i := range n {
		uniqueSet[find(i)] = struct{}{}
	}

	return len(uniqueSet)
}
