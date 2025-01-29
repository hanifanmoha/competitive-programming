package trapping_rain_water

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {

	fmt.Println("HEIGHT", height)
	n := len(height)

	total := 0
	for i := 0; i < n; i++ {
		h1 := height[i]

		if h1 == 0 {
			continue
		}

		// Search to the left (bigger)

		maxBetween := 0

		for j := i - 1; j >= 0; j-- {
			h2 := height[j]

			if h2 == h1 {
				break
			}

			if h2 > h1 {
				fmt.Printf("(%d ~> %d) %d %d : %d * %d\n", i, j, h1, h2, i-j-1, h1-maxBetween)
				total += (i - j - 1) * (h1 - maxBetween)
				break
			}

			maxBetween = max(maxBetween, h2)
		}

		// Search to the right (bigger or equal)

		maxBetween = 0

		for j := i + 1; j < n; j++ {
			h2 := height[j]

			// fmt.Println("Check", h1, h2, maxBetween)

			if h2 >= h1 {
				// fmt.Printf("(%d ~> %d) %d %d : %d * %d\n", i, j, h1, h2, j-i-1, h1-maxBetween)
				total += (j - i - 1) * (h1 - maxBetween)
				break
			}

			maxBetween = max(maxBetween, h2)
		}
	}

	return total
}
