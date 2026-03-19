package graphs

type ufrc struct {
	rep  map[int]int
	rank map[int]int
}

func (u *ufrc) find(node int) int {
	if node == u.rep[node] {
		return node
	}
	u.rep[node] = u.find(u.rep[node])
	return u.rep[node]
}

func (u *ufrc) union(nodeA, nodeB int) bool {
	repA, repB := u.find(nodeA), u.find(nodeB)
	if repA == repB {
		return false
	}

	if u.rank[repA] > u.rank[repB] {
		u.rep[repB] = repA
	} else if u.rank[repB] > u.rank[repA] {
		u.rep[repA] = repB
	} else {
		u.rep[repA] = repB
		u.rank[repB]++
	}
	return true
}

func newufrc(size int) *ufrc {
	uf := &ufrc{
		rep:  make(map[int]int),
		rank: make(map[int]int),
	}

	for i := 1; i < size+1; i++ {
		uf.rep[i] = i
		uf.rank[i] = 1
	}

	return uf
}

func findRedundantConnection(edges [][]int) []int {
	uf := newufrc(len(edges))
	for _, edge := range edges {
		if !uf.union(edge[0], edge[1]) {
			return edge
		}
	}

	return nil
}
