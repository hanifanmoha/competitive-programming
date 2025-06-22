package findmedianfromdatastream

type MedianFinder struct {
	// tree *Tree
	minHeap *Heap
	maxHeap *Heap
	tmpVal  *int
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: NewHeap(true),
		maxHeap: NewHeap(false),
	}
}

func (m *MedianFinder) AddNum(num int) {
	if m.tmpVal == nil {
		m.maxHeap.Insert(num)
		a, _ := m.maxHeap.Pop()
		m.minHeap.Insert(a)
		b, _ := m.minHeap.Pop()
		m.tmpVal = &b
	} else {
		m.maxHeap.Insert(num)
		m.maxHeap.Insert(*m.tmpVal)
		a, _ := m.maxHeap.Pop()
		m.minHeap.Insert(a)
		m.tmpVal = nil
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.tmpVal != nil {
		return float64(*m.tmpVal)
	} else {
		a, _ := m.maxHeap.Top()
		b, _ := m.minHeap.Top()
		return float64(a+b) / 2
	}
}
