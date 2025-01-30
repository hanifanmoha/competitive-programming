package minimum_number_of_arrows_to_burst_balloons

import "fmt"

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

func findMinArrowShots(points [][]int) int {

	fmt.Println("Points : ", points)

	sorted := sortIntervals(points)

	fmt.Println("Sorted : ", sorted)

	return len(points)
}
