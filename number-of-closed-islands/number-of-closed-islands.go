package number_of_closed_islands

type Islands struct {
}
type IIslands interface {
	NumberOfClosedIslands(grid [][]int) int
}

func (this Islands) NumberOfClosedIslands(grid [][]int) int {
	count, m, n := 0, len(grid), len(grid[0])
	var dfs func(i, j int) bool

	dfs = func(i, j int) bool {
		if i < 0 || j < 0 || j >= m || i >= m {
			return false
		}

		if grid[i][j] == 1 {
			return true
		}

		grid[i][j] = 1 //mark as visited

		up := dfs(i-1, j)
		down := dfs(i+1, j)
		left := dfs(i, j-1)
		right := dfs(i, j+1)
		return up && down && left && right
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				if dfs(i, j) {
					count++
				}
			}
		}
	}
	return count
}

func NewIsland() IIslands {
	return &Islands{}
}
