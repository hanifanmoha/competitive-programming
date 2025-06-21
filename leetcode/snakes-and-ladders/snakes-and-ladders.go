package snakes_and_ladders

import "fmt"

const MAXVAL = 99999

func walk(idx int, path []int, flag []bool, mem []int) int {

	n := len(path)
	if idx > n {
		return MAXVAL
	}

	if flag[idx] {
		return MAXVAL
	}

	if idx == n-1 {
		return 0
	}

	if mem[idx] != -1 {
		return mem[idx]
	}

	flag[idx] = true
	minDist := MAXVAL

	for targetIdx := idx + 1; targetIdx <= idx+6 && targetIdx < n; targetIdx++ {
		newTargetIdx := targetIdx
		if path[targetIdx] != -1 {
			newTargetIdx = path[targetIdx]
		}
		dist := walk(newTargetIdx, path, flag, mem)
		if dist < minDist {
			minDist = dist
		}
	}

	mem[idx] = minDist + 1

	flag[idx] = false

	return mem[idx]
}

func snakesAndLadders2(board [][]int) int {

	n := len(board)

	isGoingRight := true

	path := make([]int, n*n)

	for i := n - 1; i >= 0; i-- {
		for j := range n {
			r := n - 1 - i
			c := j
			idx := r*n + c
			if !isGoingRight {
				idx = r*n + n - 1 - c
			}
			path[idx] = board[i][j]
			if path[idx] != -1 {
				path[idx] -= 1
			}
		}
		isGoingRight = !isGoingRight
	}

	fmt.Println(path)

	flag := make([]bool, n*n)
	mem := make([]int, n*n)
	for i := range mem {
		mem[i] = -1
	}

	dist := walk(0, path, flag, mem)

	fmt.Println("== dist ==")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(mem[i*n+j], "\t")
		}
		fmt.Println()
	}

	if dist >= MAXVAL {
		return -1
	}
	return dist
}
