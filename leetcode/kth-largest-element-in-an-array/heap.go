package kthlargestelementinanarray

type Heap struct {
	Val []int
	Len int
}

func NewHeap() *Heap {
	return &Heap{
		Val: make([]int, 100000),
		Len: 0,
	}
}

func (h *Heap) Insert(v int) {
	h.Val[h.Len] = v

	p, c := (h.Len-1)/2, h.Len
	for c > 0 {

		if h.Val[c] > h.Val[p] {
			tmp := h.Val[c]
			h.Val[c] = h.Val[p]
			h.Val[p] = tmp
		} else {
			break
		}

		c = p
		p = (p - 1) / 2
	}

	h.Len += 1
}

func (h *Heap) Pop() int {
	mxVal := h.Val[0]
	h.Val[0] = h.Val[h.Len-1]

	p := 0
	c1, c2 := 1, 2

	for c1 < h.Len {

		c := 0
		if c2 < h.Len && h.Val[c2] > h.Val[c1] {
			c = c2
		} else {
			c = c1
		}

		if h.Val[p] < h.Val[c] {
			tmp := h.Val[p]
			h.Val[p] = h.Val[c]
			h.Val[c] = tmp
		} else {
			break
		}

		c1 = c*2 + 1
		c2 = c*2 + 2
		p = c
	}

	h.Len--

	return mxVal
}
