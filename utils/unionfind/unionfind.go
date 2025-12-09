package unionfind

type UnionFind struct {
	parent []int
	rank   []int
	size   []int
}

// Creates a new Union-Find structure with n elements
func New(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		size:   make([]int, n),
	}

	// Initially, each element is its own parent (separate set)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}

	return uf
}

// Returns the root/representative of the set containing x
// Uses path compression for optimization
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		// Path compression: make x point directly to root
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Merges the sets containing x and y
// Uses union by rank for optimization
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	// Already in same set
	if rootX == rootY {
		return false
	}

	// Union by rank: attach smaller rank tree under larger rank tree
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	} else {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
		uf.rank[rootX]++
	}

	return true
}

// Checks if x and y are in the same set
func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Returns the size of the component containing x
func (uf *UnionFind) ComponentSize(x int) int {
	return uf.size[uf.Find(x)]
}

// Returns a map of root -> slice of member indices
func (uf *UnionFind) GetComponents() map[int][]int {
	components := make(map[int][]int)

	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		components[root] = append(components[root], i)
	}

	return components
}
