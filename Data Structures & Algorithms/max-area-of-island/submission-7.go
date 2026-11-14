import(
	"slices"
)

func maxAreaOfIsland(grid [][]int) int {
    visited := make(map[[2]int]struct{})

	groupid := 0

	groups := []int{}
	directions := [][2]int{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}
	for i := range grid{
		for j := range grid[i]{
			que := [][2]int{}
			if _, seen := visited[[2]int{i,j}]; seen{
				continue
			}
			if grid[i][j] == 0{
				continue
			}
			
			groupid ++ 
			var counter = 1
			visited[[2]int{i, j}] = struct{}{}
			que = append(que, [2]int{i,j})
			for len(que)>0{
				current := que[0]
				row,col := current[0], current[1]
				que = que[1:]
				for _, direction:= range directions{
					nextrow := row + direction[0]
					nextcol := col + direction[1]

					if nextrow < 0 || nextrow >= len(grid) {
    					continue
					}
					if nextcol < 0 || nextcol >= len(grid[nextrow]) {
						continue
					}

					if _, seen := visited[[2]int{nextrow,nextcol}]; seen{
						continue
					}
					if grid[nextrow][nextcol] == 0{
						continue
					}

					visited[[2]int{nextrow, nextcol}] = struct{}{}
					que = append(que, [2]int{nextrow,nextcol})
					counter+=1
				}

			}

			groups = append(groups, counter)
		}
	}
	if len(groups)==0{
		return 0
	}
	largest := slices.Max(groups)
	return largest 
}
