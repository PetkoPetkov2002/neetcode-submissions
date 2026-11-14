func canFinish(numCourses int, prerequisites [][]int) bool {
    inPath := make([]bool, numCourses)
	visited := make([]bool,numCourses) 
	graph := make([][]int, numCourses)

	for _, pair := range prerequisites {
		course := pair[0]
		prerequisite := pair[1]

		graph[prerequisite] = append(graph[prerequisite], course)
	}


	var dfs func(int) bool

	dfs = func(node int) bool{
		if inPath[node]{
			return true
		}

		if visited[node]{
			return false
		}

		visited[node] = true
		inPath[node] = true

		for _, value := range graph[node]{
			if dfs(value){
				return true
			}
			
		}
		inPath[node] = false
		return false
	}
	for i:=0; i<numCourses; i++{
		if !visited[i]{
			if dfs(i){
				return false
			}
		}
	}
	return true
}
