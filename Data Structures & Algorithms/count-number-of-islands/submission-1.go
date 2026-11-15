func numIslands(grid [][]byte) int {
    visited := make(map[[2]int]bool)
	islands := 0
	directions := [][2]int{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
    }

	var bfs func(coordinates [2]int)
	bfs = func(coordinates [2]int){
		q := [][2]int{}
		visited[[2]int{coordinates[0],coordinates[1]}] = true
		q = append(q, [2]int{coordinates[0],coordinates[1]})
		for len(q) > 0{
			current := q[0]
			q = q[1:]
			for _, directions := range directions{
				r := current[0] + directions[0]
				c := current[1] + directions[1]
				if r < 0 || r >= len(grid) ||c < 0 || c >= len(grid[r]){
            		continue
        		}
				if grid[r][c] == '1' && !visited[[2]int{r,c}]{
					q = append(q, [2]int{r,c})
					visited[[2]int{r,c}] = true
				}
			}
		}

	}
	for i, row := range grid{
		for j := range row{
			if !visited[[2]int{i,j}] && grid[i][j] == '1'{
				bfs([2]int{i,j})
				islands ++
			}

		}
	}
	return islands
}
