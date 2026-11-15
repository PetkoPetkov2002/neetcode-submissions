func countComponents(n int, edges [][]int) int {
    parent := make([]int, n)
	size := make([]int, n)
	components := n
	for i:=0; i < n; i++{
		parent[i] = i
		size[i] = 1
	}
	var union func(rootA, rootB int)
	union = func(rootA, rootB int){
		if size[rootA]>size[rootB]{
			parent[rootB] = parent[rootA]
			size[rootA] += size[rootB]

		}else{
			parent[rootA] = parent[rootB]
			size[rootB] = size[rootA]
		}
	}
	var find func(nodeA int) int
	find = func(node int) int{
		if parent[node] != node{
			root := find(parent[node])
			parent[node] = root
		}
		return parent[node]
	}
	for _,edge := range edges{
		rootA := find(edge[0])
		rootB := find(edge[1])
		if rootA != rootB{
			union(rootA, rootB)
			components --
		}
	}
	return components

}


