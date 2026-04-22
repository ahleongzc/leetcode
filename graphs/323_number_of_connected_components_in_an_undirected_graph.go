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

type countComponentsUF struct {
	rank map[int]int
	rep  map[int]int
}

func (c *countComponentsUF) find(node int) int {
	if c.rep[node] == node {
		return node
	}
	c.rep[node] = c.find(c.rep[node])
	return c.rep[node]
}

func (c *countComponentsUF) union(nodeA, nodeB int) bool {
	repA, repB := c.find(nodeA), c.find(nodeB)
	if repA == repB {
		return false
	}

	if c.rank[repA] > c.rank[repB] {
		c.rep[repB] = repA
	} else if c.rank[repB] > c.rank[repA] {
		c.rep[repA] = repB
	} else {
		c.rep[repA] = repB
		c.rank[repB]++
	}

	return true
}

func newCountComponentsUF(n int) *countComponentsUF {
	uf := &countComponentsUF{
		rep:  make(map[int]int),
		rank: make(map[int]int),
	}

	for i := range n {
		uf.rep[i] = i
		uf.rank[i] = 0
	}

	return uf
}

func countComponentsUnionFind(n int, edges [][]int) int {
	uf := newCountComponentsUF(n)
	for _, edge := range edges {
		uf.union(edge[0], edge[1])
	}

	uniqueRep := make(map[int]struct{})
	for i := range n {
		uniqueRep[uf.find(i)] = struct{}{}
	}

	return len(uniqueRep)
}
