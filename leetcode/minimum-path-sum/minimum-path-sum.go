package minimumpathsum

const MAX_VALUE = 99999999

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {

	m, n := len(grid), len(grid[0])

	for i := range m {
		for j := range n {

			if i == 0 && j == 0 {
				continue
			}

			top, left := MAX_VALUE, MAX_VALUE
			if i > 0 {
				top = grid[i-1][j]
			}
			if j > 0 {
				left = grid[i][j-1]
			}

			grid[i][j] = min(top, left) + grid[i][j]
		}
	}

	return grid[m-1][n-1]
}
