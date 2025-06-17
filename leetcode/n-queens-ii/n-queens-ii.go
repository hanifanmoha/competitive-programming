package nqueensii

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isSameDiagonal(x1, y1, x2, y2 int) bool {
	return abs(x1-x2) == abs(y1-y2)
}

func iterate(N int, queens [][]int, row int, res *int) {

	if row >= N {
		*res++
		return
	}

	for col := range N {

		isIntersect := false
		for q := range row {
			qx := queens[q][0]
			qy := queens[q][1]

			if qx == col || isSameDiagonal(qx, qy, col, row) {
				isIntersect = true
				break
			}
		}
		if !isIntersect {
			// fmt.Println("=>", col, ",", row, " queens : ", queens)
			queens[row] = []int{col, row}
			iterate(N, queens, row+1, res)
		}
	}
}

func totalNQueens(N int) int {
	queens := make([][]int, N)
	res := 0
	iterate(N, queens, 0, &res)
	return res
}
