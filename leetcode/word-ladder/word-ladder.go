package wordladder

type Node struct {
	word      string
	flag      bool
	mutations []*Node
}

func NewNode(word string) *Node {
	return &Node{
		word:      word,
		mutations: make([]*Node, 0),
	}
}

func (n *Node) AddMutation(m *Node) {
	n.mutations = append(n.mutations, m)
}

type QueueItem struct {
	node *Node
	step int
}

func isNeighbor(wordA, wordB string) bool {
	count := 0
	for i := range len(wordA) {
		if wordA[i] != wordB[i] {
			count++
		}
	}
	return count == 1
}

func minMutation(beginWord string, endWord string, bank []string) int {

	N := len(bank)
	validWords := make(map[string]*Node)
	for i := range N {
		word := bank[i]
		validWords[word] = NewNode(word)
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			wordA, wordB := bank[i], bank[j]
			vga, vgb := validWords[wordA], validWords[wordB]
			if isNeighbor(wordA, wordB) {
				vga.AddMutation(vgb)
				vgb.AddMutation(vga)
			}
		}
	}

	// calculate step

	queue := make([]*QueueItem, 0)
	idx := 0

	// if start word valid, add it as initial queue

	if vg, exist := validWords[beginWord]; exist {
		vg.flag = true
		queue = append(queue, &QueueItem{node: vg, step: 1})
	} else {
		// if start word invalid, add its mutation as initial queue
		for _, vg := range validWords {
			if isNeighbor(vg.word, beginWord) {
				vg.flag = true
				queue = append(queue, &QueueItem{node: vg, step: 2})
			}
		}
	}

	for idx < len(queue) {

		qItem := queue[idx]
		qNode := qItem.node
		qStep := qItem.step

		if qNode.word == endWord {
			return qStep
		}

		// add neighbor to queue if not yet added
		for _, m := range qNode.mutations {
			if !m.flag {
				m.flag = true
				queue = append(queue, &QueueItem{
					node: m,
					step: qStep + 1,
				})
			}
		}

		idx++
	}

	return 0
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	return minMutation(beginWord, endWord, wordList)
}
