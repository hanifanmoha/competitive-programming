package game_of_life

// Any live cell with fewer than two live neighbors dies as if caused by under-population.
// Any live cell with two or three live neighbors lives on to the next generation.
// Any live cell with more than three live neighbors dies, as if by over-population.
// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

func deepCopy(board [][]int) [][]int {
	m, n := len(board), len(board[0])
	res := make([][]int, m)
	for i := range m {
		res[i] = make([]int, n)
		for j := range n {
			res[i][j] = board[i][j]
		}
	}
	return res
}

func getCell(board [][]int, x, y int) int {
	m, n := len(board), len(board[0])
	if x < 0 || x >= n || y < 0 || y >= m {
		return 0
	}
	return board[y][x]
}

func getNeighbor(board [][]int, x, y int) int {
	cnt := 0
	cnt += getCell(board, x-1, y-1)
	cnt += getCell(board, x, y-1)
	cnt += getCell(board, x+1, y-1)
	cnt += getCell(board, x-1, y)
	cnt += getCell(board, x+1, y)
	cnt += getCell(board, x-1, y+1)
	cnt += getCell(board, x, y+1)
	cnt += getCell(board, x+1, y+1)
	return cnt
}

func isLive(curr int, neighbor int) int {
	if neighbor < 2 || neighbor > 3 {
		return 0
	}
	if neighbor == 3 {
		return 1
	}
	return curr
}

func gameOfLife(board [][]int) {

	original := deepCopy(board)

	m, n := len(board), len(board[0])

	// for y := range m {
	// 	fmt.Println(original[y])
	// }
	// fmt.Println("=========")

	for y := range m {
		for x := range n {
			neighbor := getNeighbor(original, x, y)
			// fmt.Printf("(%d, %d) : %d\n", x, y, neighbor)
			board[y][x] = isLive(original[y][x], neighbor)
		}
	}

}
