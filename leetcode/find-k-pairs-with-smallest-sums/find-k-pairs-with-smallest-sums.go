package findkpairswithsmallestsums

type HeapNode struct {
	p1, p2 int
	sum    int
}

type Heap struct {
	data []*HeapNode
}

func NewHeap() *Heap {
	return &Heap{
		data: make([]*HeapNode, 0),
	}
}

func (h *Heap) Insert(p1, p2, sum int) {
	node := HeapNode{
		p1:  p1,
		p2:  p2,
		sum: sum,
	}

	h.data = append(h.data, &node)
	h.up(len(h.data) - 1)
}

func (h *Heap) Pop() (int, int, int) {
	res := h.data[0]

	n := len(h.data)
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]

	h.down(0)

	return res.p1, res.p2, res.sum
}

func (h *Heap) up(c int) {
	p := (c - 1) / 2
	if c == 0 {
		return
	}

	if h.data[p].sum < h.data[c].sum {
		return
	}

	h.data[c], h.data[p] = h.data[p], h.data[c]

	h.up(p)
}

func (h *Heap) down(p int) {
	c1, c2 := (p*2 + 1), (p*2 + 2)

	n := len(h.data)
	// switchMode := 0 // 0 : no switch, 1 : switch left, 2 : switch right

	if c2 < n {

		// switch to c2
		if h.data[c2].sum < h.data[p].sum && h.data[c2].sum < h.data[c1].sum {
			h.data[p], h.data[c2] = h.data[c2], h.data[p]
			h.down(c2)

			// switch to c1
		} else if h.data[c1].sum < h.data[p].sum {
			h.data[p], h.data[c1] = h.data[c1], h.data[p]
			h.down(c1)
		}

	} else if c1 < n {

		// switch to c1
		if h.data[c1].sum < h.data[p].sum {
			h.data[p], h.data[c1] = h.data[c1], h.data[p]
			h.down(c1)
		}

	}

}

func mark(flag map[int]map[int]bool, p1, p2 int) {
	if flag[p1] == nil {
		flag[p1] = make(map[int]bool)
	}
	flag[p1][p2] = true
}

func isMarked(flag map[int]map[int]bool, p1, p2 int) bool {
	if flag[p1] == nil {
		return false
	}
	return flag[p1][p2]
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	n1, n2 := len(nums1), len(nums2)

	flag := make(map[int]map[int]bool)

	h := NewHeap()

	mark(flag, 0, 0)
	h.Insert(0, 0, nums1[0]+nums2[0])

	res := make([][]int, 0)

	for range k {
		p1, p2, _ := h.Pop()
		res = append(res, []int{nums1[p1], nums2[p2]})

		a1, a2 := p1+1, p2
		if a1 >= n1 {
			a1 = 0
			a2 = p2 + 1
		}

		b1, b2 := p1, p2+1
		if b2 >= n2 {
			b1 = p1 + 1
			b2 = 0
		}

		if a1 < n1 && a2 < n2 && !isMarked(flag, a1, a2) {
			mark(flag, a1, a2)
			h.Insert(a1, a2, nums1[a1]+nums2[a2])
		}

		if b1 < n1 && b2 < n2 && !isMarked(flag, b1, b2) {
			mark(flag, b1, b2)
			h.Insert(b1, b2, nums1[b1]+nums2[b2])
		}
	}

	return res
}
