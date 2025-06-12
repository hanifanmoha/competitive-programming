package designaddandsearchwordsdatastructure

type Node struct {
	c      rune
	isLeaf bool
	next   map[rune]*Node
}

func NewNode(c rune) *Node {
	return &Node{
		c:    c,
		next: make(map[rune]*Node),
	}
}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: NewNode('#'),
	}
}

func (w *WordDictionary) AddWord(word string) {
	parent := w.root
	for i, c := range word {
		next, exist := parent.next[c]
		if !exist {
			next = NewNode(c)
			next.isLeaf = i == len(word)-1
			parent.next[c] = next
		} else if i == len(word)-1 {
			next.isLeaf = true
		}
		parent = next
	}
}

func (w *WordDictionary) Search(word string) bool {
	return searchWord(word, 0, w.root)
}

func searchWord(word string, idx int, parent *Node) bool {
	c := rune(word[idx])

	if c != '.' {
		next, exist := parent.next[c]
		if !exist {
			return false
		}

		if idx == len(word)-1 {
			return next.isLeaf
		}

		return searchWord(word, idx+1, next)
	}

	for _, next := range parent.next {

		if idx == len(word)-1 && next.isLeaf {
			return true
		} else if idx < len(word)-1 && searchWord(word, idx+1, next) {
			return true
		}

	}

	return false

}
