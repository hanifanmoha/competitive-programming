package merge_intervals

func sortIntervals(intervals [][]int) [][]int {

	n := len(intervals)
	mid := n / 2

	if n <= 1 {
		return intervals
	}

	a := sortIntervals(intervals[:mid])
	b := sortIntervals(intervals[mid:])

	// merge
	ia, ib := 0, 0
	na, nb := len(a), len(b)

	merged := make([][]int, 0)
	for ia < na || ib < nb {
		if ia >= na {
			merged = append(merged, b[ib])
			ib++
		} else if ib >= nb {
			merged = append(merged, a[ia])
			ia++
		} else if b[ib][0] < a[ia][0] {
			merged = append(merged, b[ib])
			ib++
		} else {
			merged = append(merged, a[ia])
			ia++
		}
	}

	return merged
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func checkIfSorted(intervals [][]int) bool {
	n := len(intervals)

	if n == 0 {
		return true
	}

	for i := 1; i < n; i++ {
		if intervals[i][0] < intervals[i-1][0] {
			return false
		}
	}

	return true
}

func merge2(intervals [][]int) [][]int {
	newIntervals := make([][]int, 0)

	for _, interval := range intervals {

		ni := len(newIntervals)

		// check if not merge
		if ni == 0 {
			newIntervals = append(newIntervals, interval)
			continue
		}

		last := newIntervals[ni-1]
		if interval[0] > last[1] {
			newIntervals = append(newIntervals, interval)
			continue
		}

		// if merge
		last[0] = min(last[0], interval[0])
		last[1] = max(last[1], interval[1])

	}

	return newIntervals
}

func merge(intervals [][]int) [][]int {
	sorted := intervals
	if !checkIfSorted(sorted) {
		sorted = sortIntervals(intervals)
	}
	return merge2(sorted)
}
