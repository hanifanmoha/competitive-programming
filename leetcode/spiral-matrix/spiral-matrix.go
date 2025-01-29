package spiral_matrix

const RIGHT = 1001
const LEFT = 1002
const UP = 1003
const DOWN = 1004

func isNextValid(flag [][]bool, x, y int) bool {
	r, c := len(flag), len(flag[0])

	if x < 0 || x >= c || y < 0 || y >= r {
		return false
	}

	if flag[y][x] {
		return false
	}

	return true

}

func iterate(res []int, matrix [][]int, flag [][]bool, x, y, mode int) []int {

	// fmt.Printf("(%d, %d)\n", x, y)

	res = append(res, matrix[y][x])
	flag[y][x] = true

	// NORMAL DIRECTION
	x2, y2 := x+1, y
	if mode == LEFT {
		x2, y2 = x-1, y
	} else if mode == DOWN {
		x2, y2 = x, y+1
	} else if mode == UP {
		x2, y2 = x, y-1
	}

	// CHANGE DIRECTION
	if !isNextValid(flag, x2, y2) {
		if mode == RIGHT {
			x2, y2 = x, y+1
			mode = DOWN
		} else if mode == LEFT {
			x2, y2 = x, y-1
			mode = UP
		} else if mode == DOWN {
			x2, y2 = x-1, y
			mode = LEFT
		} else if mode == UP {
			x2, y2 = x+1, y
			mode = RIGHT
		}
	}

	if !isNextValid(flag, x2, y2) {
		return res
	}

	return iterate(res, matrix, flag, x2, y2, mode)
}

func spiralOrder(matrix [][]int) []int {
	r, c := len(matrix), len(matrix[0])

	flag := make([][]bool, r)
	for i := range flag {
		flag[i] = make([]bool, c)
	}

	res := make([]int, 0)

	newRes := iterate(res, matrix, flag, 0, 0, RIGHT)

	return newRes
}
