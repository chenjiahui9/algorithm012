func numIslands(grid [][]byte) int {

	count := 0
	for line := range grid {
		for col := range grid[line] {
			if grid[line][col] == '0'  {
				continue
			}

			count++
			dfs(&grid,line,col)
		}
	}

	return count
}

func dfs(grid *[][]byte,i int,j int) {

	if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[0]) {
		return
	}

	if (*grid)[i][j] != '1' {
		return
	}

	(*grid)[i][j] = '0'

	dfs(grid,i+1,j)
	dfs(grid,i-1,j)
	dfs(grid,i,j-1)
	dfs(grid,i,j+1)

}