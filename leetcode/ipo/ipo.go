package ipo

import "fmt"

type Project struct {
	profit  int
	capital int
}

type Heap struct {
	data []*Project
}

func NewHeap() *Heap {
	return &Heap{
		data: make([]*Project, 0),
	}
}

func (h *Heap) Insert(profit, capital int) {
	node := Project{profit, capital}

	h.data = append(h.data, &node)
	h.up(len(h.data) - 1)
}

func (h *Heap) Pop() *Project {

	if len(h.data) == 0 {
		return nil
	}

	res := h.data[0]

	n := len(h.data)
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]

	h.down(0)

	return res
}

func (h *Heap) up(c int) {
	p := (c - 1) / 2
	if c == 0 {
		return
	}

	if h.data[p].profit > h.data[c].profit {
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
		if h.data[c2].profit > h.data[p].profit && h.data[c2].profit > h.data[c1].profit {
			h.data[p], h.data[c2] = h.data[c2], h.data[p]
			h.down(c2)

			// switch to c1
		} else if h.data[c1].profit > h.data[p].profit {
			h.data[p], h.data[c1] = h.data[c1], h.data[p]
			h.down(c1)
		}

	} else if c1 < n {

		// switch to c1
		if h.data[c1].profit > h.data[p].profit {
			h.data[p], h.data[c1] = h.data[c1], h.data[p]
			h.down(c1)
		}

	}

}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {

	h := NewHeap()
	n := len(profits)

	for i := range n {
		h.Insert(profits[i], capital[i])
	}

	for i := range k {
		var project *Project
		stashedProjects := make([]Project, 0)
		for {
			project = h.Pop()
			if project == nil {
				break
			} else if project.capital <= w {
				break
			} else {
				stashedProjects = append(stashedProjects, *project)
			}
		}
		for i := range stashedProjects {
			h.Insert(stashedProjects[i].profit, stashedProjects[i].capital)
		}
		if project == nil {
			break
		}
		fmt.Println(i, project.profit, project.capital, w)
		w += project.profit
	}

	return w
}
