package wordsearchii

type Trie struct {
	word *string
	next []*Trie
}

func NewTrie() *Trie {
	return &Trie{
		word: nil,
		next: make([]*Trie, 26),
	}
}

func buildTree(words []string) *Trie {
	root := NewTrie()
	for _, word := range words {
		parent := root
		for i := range word {
			idx := int(word[i] - 'a')
			if parent.next[idx] == nil {
				parent.next[idx] = NewTrie()
			}
			parent = parent.next[idx]
		}
		parent.word = &word
	}
	return root
}

func traverseWord(root *Trie, res *[]string) {
	if root == nil {
		return
	}
	if root.word != nil {
		// fmt.Println(*root.word)
		*res = append(*res, *root.word)
	}
	for i := range root.next {
		traverseWord(root.next[i], res)
	}
}

func search(board [][]byte, parent *Trie, row, col int, res *[]string) {

	m, n := len(board), len(board[0])
	if row < 0 || row >= m {
		return
	}
	if col < 0 || col >= n {
		return
	}
	if board[row][col] == '#' {
		return
	}

	tmpChar := board[row][col]
	idx := int(tmpChar - 'a')
	current := parent.next[idx]
	if current == nil {
		return
	}

	if current.word != nil {
		*res = append(*res, *current.word)
		current.word = nil
	}

	board[row][col] = '#'
	search(board, current, row-1, col, res)
	search(board, current, row+1, col, res)
	search(board, current, row, col-1, res)
	search(board, current, row, col+1, res)
	board[row][col] = tmpChar
}

func findWords(board [][]byte, words []string) []string {

	// N := 26

	m, n := len(board), len(board[0])

	root := buildTree(words)
	res := make([]string, 0)
	// traverseWord(root, &res)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			search(board, root, i, j, &res)
		}
	}

	return res
}
