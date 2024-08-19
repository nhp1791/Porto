package models

type neighborhood struct {
	neighbors []*neighbor
}

func newNeighborhood() *neighborhood {
	return &neighborhood{
		neighbors: make([]*neighbor, MaxNearestNeighbors),
	}
}

// insert simply creates a nearest neighbor set as each
// load is considered, rather than sorting a map later.
// This is simply one solution to creating a set of nearest neighbors.
func (n *neighborhood) insert(dist uint64, load *Load) {
	for i := 0; i < MaxNearestNeighbors; i++ {
		if n.neighbors[i] == nil {
			n.neighbors[i] = &neighbor{
				load: load,
				dist: dist,
			}
			return
		}
		if dist <= n.neighbors[i].dist {
			for j := MaxNearestNeighbors - 2; j >= i; j-- {
				n.neighbors[j+1] = n.neighbors[j]
			}
			n.neighbors[i] = &neighbor{
				load: load,
				dist: dist,
			}
			return
		}
	}
}

// prune just removes any neighbor slots that were not filled
func (n *neighborhood) prune() {
	neighbors := []*neighbor{}
	for i := 0; i < MaxNearestNeighbors; i++ {
		if n.neighbors[i] != nil {
			neighbors = append(neighbors, n.neighbors[i])
		}
	}
	n.neighbors = neighbors
}

type neighbor struct {
	load *Load
	dist uint64
}
