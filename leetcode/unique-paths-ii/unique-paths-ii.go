package unique_paths_ii

func rec(grid [][]int, mem [][]int, x, y int) int {
	r, c := len(grid), len(grid[0])

	if y == r-1 && x == c-1 {
		return 1
	}

	if x >= c || y >= r {
		return 0
	}

	if grid[y][x] == 1 {
		return 0
	}

	if mem[y][x] >= 0 {
		return mem[y][x]
	}

	toRight := rec(grid, mem, x+1, y)
	toBottom := rec(grid, mem, x, y+1)

	mem[y][x] = toRight + toBottom
	return toRight + toBottom

}

func uniquePathsWithObstacles(grid [][]int) int {

	r, c := len(grid), len(grid[0])
	mem := make([][]int, r)
	for i := range mem {
		mem[i] = make([]int, c)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	return rec(grid, mem, 0, 0)
}
