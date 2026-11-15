func topologicalSort(n int, edges [][]int) []int {
	inPath := make([]bool, n)
	visited := make([]bool,n) 
	graph := make([][]int, n)
	solution := []int{}

	for _, pair := range edges {
		node := pair[0]
		connection := pair[1]

		graph[node] = append(graph[node], connection)
	}


	var dfs func(node int) bool
	dfs = func(node int) bool {
		if inPath[node]{
			return true
		}
		
		if visited[node]{
			return false
		}
		
		visited[node] = true
		inPath[node] = true

		for _, neighbour := range graph[node]{
			cycle := dfs(neighbour)
			if cycle{
				return true
			}
		}
		inPath[node] = false
		solution = append(solution, node)
		return false
	}
	for i := range graph{
		if !visited[i]{
			cycle := dfs(i)
			if cycle{
				null := []int{}
				return null
			}
		}
	}
	for i, j := 0, len(solution)-1; i < j; i, j = i+1, j-1 {
        solution[i], solution[j] = solution[j], solution[i]
    }
	return solution

}
