package maxpointsonaline

type Point struct {
	x int
	y int
}

func isSameDirection(a, b Point) bool {
	return a.x*b.y == a.y*b.x
	// return a.x > a.y
}

func maxPoints(points_ [][]int) int {

	N := len(points_)

	if N == 1 {
		return 1
	}

	if N == 2 {
		return 2
	}

	points := make([]Point, N)
	for i := range N {
		points[i] = Point{
			x: points_[i][0],
			y: points_[i][1],
		}
	}

	maxCount := 0

	for i := range N {

		for j := range N {

			count := 2

			vectorA := Point{
				x: points[i].x - points[j].x,
				y: points[i].y - points[j].y,
			}

			for k := j + 1; k < N; k++ {
				if i == j || i == k {
					continue
				}

				vectorB := Point{
					x: points[i].x - points[k].x,
					y: points[i].y - points[k].y,
				}

				// fmt.Printf("%d->%d; %d->%d :: (%d,%d) (%d,%d) :: %t \n", i, j, i, k, vectorA.x, vectorA.y, vectorB.x, vectorB.y, isSameDirection(vectorA, vectorB))

				if isSameDirection(vectorA, vectorB) {
					count++
				}

			}

			if count > maxCount {
				maxCount = count
			}

		}
	}

	return maxCount
}
