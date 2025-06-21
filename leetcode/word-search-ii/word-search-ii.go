package wordsearchii

func findWord(board [][]byte, flag [][]bool, word string, idx, r, c int) bool {
	if idx >= len(word) {
		return true
	}

	y, x := len(board), len(board[0])

	if r < 0 || c < 0 || r >= y || c >= x {
		return false
	}

	if flag[r][c] {
		return false
	}

	if word[idx] != board[r][c] {
		return false
	}

	flag[r][c] = true
	res := findWord(board, flag, word, idx+1, r, c+1) ||
		findWord(board, flag, word, idx+1, r, c-1) ||
		findWord(board, flag, word, idx+1, r+1, c) ||
		findWord(board, flag, word, idx+1, r-1, c)
	flag[r][c] = false

	return res
}

func initFindWord(board [][]byte, flag [][]bool, word string) bool {
	y, x := len(board), len(board[0])
	for r := range y {
		for c := range x {
			if findWord(board, flag, word, 0, r, c) {
				return true
			}
		}
	}
	return false
}

func findWords2(board [][]byte, words []string) []string {

	y, x := len(board), len(board[0])
	flag := make([][]bool, y)
	for i := range y {
		flag[i] = make([]bool, x)
	}

	res := make([]string, 0)
	for _, word := range words {
		if initFindWord(board, flag, word) {
			res = append(res, word)
		}
	}

	return res
}
