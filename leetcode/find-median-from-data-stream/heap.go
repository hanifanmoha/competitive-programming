package findmedianfromdatastream

type Heap struct {
	minHeap bool
	data    []int
}

func NewHeap(isMinHeap bool) *Heap {
	return &Heap{
		minHeap: isMinHeap,
		data:    make([]int, 0),
	}
}

func (h *Heap) Insert(val int) {
	h.data = append(h.data, val)
	h.up(len(h.data) - 1)
}

func (h *Heap) Top() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

func (h *Heap) Pop() (int, bool) {

	if len(h.data) == 0 {
		return 0, false
	}

	res := h.data[0]

	n := len(h.data)
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]

	h.down(0)

	return res, true
}

func (h *Heap) isNeedSwitch(a, b int) bool {
	if h.minHeap {
		return a < b
	} else {
		return a > b
	}
}

func (h *Heap) up(c int) {
	p := (c - 1) / 2
	if c == 0 {
		return
	}

	if !h.isNeedSwitch(h.data[c], h.data[p]) {
		return
	}

	h.data[c], h.data[p] = h.data[p], h.data[c]

	h.up(p)
}

func (h *Heap) down(p int) {
	c1, c2 := (p*2 + 1), (p*2 + 2)

	n := len(h.data)

	if c2 < n {

		// switch to c2
		if h.isNeedSwitch(h.data[c2], h.data[p]) && h.isNeedSwitch(h.data[c2], h.data[c1]) {
			h.data[p], h.data[c2] = h.data[c2], h.data[p]
			h.down(c2)

			// switch to c1
		} else if h.isNeedSwitch(h.data[c1], h.data[p]) {
			h.data[p], h.data[c1] = h.data[c1], h.data[p]
			h.down(c1)
		}

	} else if c1 < n {

		// switch to c1
		if h.isNeedSwitch(h.data[c1], h.data[p]) {
			h.data[p], h.data[c1] = h.data[c1], h.data[p]
			h.down(c1)
		}

	}

}
