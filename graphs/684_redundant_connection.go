package graphs

func FindRedundantConnection(edges [][]int) []int {
	parent := make(map[int]int)
	rank := make(map[int]int)

	for i := range len(edges) {
		parent[i] = i
		rank[i] = 0
	}

	find := func(i int) int {
		p := parent[i]
		for {
			if p == parent[p] {
				break
			}
			parent[i] = parent[parent[i]]
			p = parent[i]
		}
		return p
	}

	union := func(i, j int) bool {
		p1, p2 := find(i), find(j)
		if p1 == p2 {
			return false
		}

		if rank[p1] > rank[p2] {
			parent[p2] = p1
		} else {
			parent[p1] = p2
			rank[p1]++
		}

		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}

	return nil
}
