package graphs

type unionFindFindRedundantConnection struct {
	rep  map[int]int
	rank map[int]int
}

func newUnionFindFindRedundantConnection(numNode int) *unionFindFindRedundantConnection {
	uf := &unionFindFindRedundantConnection{
		rep:  make(map[int]int),
		rank: make(map[int]int),
	}

	for i := range numNode {
		uf.rep[i+1] = i + 1
		uf.rank[i+1] = 0
	}
	return uf
}

func (uf *unionFindFindRedundantConnection) find(node int) int {
	if uf.rep[node] == node {
		return node
	}
	uf.rep[node] = uf.find(uf.rep[node])
	return uf.rep[node]
}

func (uf *unionFindFindRedundantConnection) union(nodeA, nodeB int) bool {
	repA, repB := uf.find(nodeA), uf.find(nodeB)
	if repA == repB {
		return false
	}

	rankRepA, rankRepB := uf.rank[repA], uf.rank[repB]
	if rankRepA > rankRepB {
		uf.rep[repB] = repA
	} else if rankRepB > rankRepA {
		uf.rep[repA] = repB
	} else {
		uf.rep[repA] = repB
		uf.rank[repB]++
	}
	return true
}

func findRedundantConnection(edges [][]int) []int {
	uf := newUnionFindFindRedundantConnection(len(edges))

	for _, edge := range edges {
		if !uf.union(edge[0], edge[1]) {
			return []int{
				edge[0], edge[1],
			}
		}
	}
	return nil
}
