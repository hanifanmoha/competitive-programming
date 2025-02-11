package minimum_number_of_arrows_to_burst_balloons

import "fmt"

// sort using end
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
		} else if b[ib][1] < a[ia][1] {
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

	end := sorted[0][1]
	cnt := 1
	for _, point := range sorted[1:] {
		if point[0] > end {
			end = point[1]
			cnt++
		}
	}

	return cnt
}
