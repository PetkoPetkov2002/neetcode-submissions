func shortestPath(grid [][]int) int {
	if grid[0][0] == 1 || grid[len(grid)-1][len(grid[0])-1] == 1 {
		return -1
	}
	directions := [][2]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}
	queue := [][3]int{}
	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)

	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	
	queue = append(queue, [3]int{0,0,0})
	visited[0][0] = true
	for len(queue)>0{
		element := queue[0]
		queue = queue[1:]
		currRow := element[0]
		currCol := element[1]
		distance := element[2]
		if currRow == rows-1 && currCol == cols-1{
			return element[2]
		}
		for _, direction := range directions{
			r := currRow + direction[0]
			c := currCol + direction[1]
			if r < 0 || r >= rows || c < 0 || c >= cols{
				continue
			}
			if visited[r][c] || grid[r][c] == 1{
				continue
			}
			if visited[r][c] != true{
				queue = append(queue, [3]int{r, c, distance+1})
				visited[r][c] = true
			}
		}
	}

	return -1

}
