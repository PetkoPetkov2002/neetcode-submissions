type UnionFind struct {
  Root []int
  Size []int
  Components int
}

func NewUnionFind(n int) *UnionFind {
	size := make([]int, n)
	root := make([]int, n)
	for i := range size{
		size[i] = 1
		root[i] = i 
	}
	

	structure := UnionFind{Root: root, Size: size, Components: n}
	return &structure
}

func (uf *UnionFind) Find(x int) int {
	if uf.Root[x] != x {
		uf.Root[x] = uf.Find(uf.Root[x])

	}
	return uf.Root[x]
}

func (uf *UnionFind) IsSameComponent(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY{
		return false
	}else{
		return true
	}
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY{
		return false
	}
	if uf.Size[rootX] > uf.Size[rootY]{
		uf.Root[rootY] = rootX
		uf.Size[rootX] += uf.Size[rootY]
	}else{
		uf.Root[rootX] = rootY
		uf.Size[rootY] += uf.Size[rootX]
	}
	uf.Components --
	return true
}

func (uf *UnionFind) GetNumComponents() int {
	return uf.Components
}
