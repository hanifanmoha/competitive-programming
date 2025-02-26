package surroundedregions

func traverse(board [][]byte, y, x int) {
	r, c := len(board), len(board[0])
	if x < 0 || x >= c {
		return
	}
	if y < 0 || y >= r {
		return
	}
	if board[y][x] == 'X' {
		return
	}
	if board[y][x] == 'O' {
		return
	}
	board[y][x] = 'O'
	traverse(board, y-1, x)
	traverse(board, y+1, x)
	traverse(board, y, x-1)
	traverse(board, y, x+1)
}

func solve(board [][]byte) {
	r, c := len(board), len(board[0])
	for y, row := range board {
		for x, cell := range row {
			if cell == 'O' {
				board[y][x] = '-'
			}
		}
	}
	for i := 0; i < r; i++ {
		traverse(board, i, 0)
		traverse(board, i, c-1)
	}
	for i := 0; i < c; i++ {
		traverse(board, 0, i)
		traverse(board, r-1, i)
	}
	for y, row := range board {
		for x, cell := range row {
			if cell == '-' {
				board[y][x] = 'X'
			}
		}
	}
}
