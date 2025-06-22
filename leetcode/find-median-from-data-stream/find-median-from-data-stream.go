package findmedianfromdatastream

type MedianFinder struct {
	tree *Tree
}

func Constructor() MedianFinder {
	return MedianFinder{tree: &Tree{}}
}

func (m *MedianFinder) AddNum(num int) {
	// fmt.Println("AddNum", num)
	m.tree.Insert(num)
}

func (m *MedianFinder) FindMedian() float64 {
	med := m.tree.FindMedian()
	// fmt.Println("FindMedian", med)
	return med
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
