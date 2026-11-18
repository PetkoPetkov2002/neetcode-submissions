func countPaths(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
    visited := make([][]bool, rows)
	count := 0

    for i := range visited {
        visited[i] = make([]bool, cols)
    }
	directions := [][2]int{
        {-1, 0}, // Up
        {1, 0},  // Down
        {0, -1}, // Left
        {0, 1},  // Right
    }

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
            return
        }
		if grid[r][c] == 1 || visited[r][c] {
            return
        }
		if visited[r][c]{
			return
		}
		if r == rows-1 && c == cols-1 {
            count++
            return
        }
		visited[r][c] = true

		for _, d := range directions{
			dfs(r+d[0], c+d[1])
		}

		visited[r][c] = false

	}
	dfs(0,0)
    return count
}
