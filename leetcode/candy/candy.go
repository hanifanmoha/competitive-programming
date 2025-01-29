package candy

import (
	"fmt"
)

func candy(ratings []int) int {

	candies := make([]int, len(ratings))
	candies[0] = 1

	for i := range ratings {
		if i == 0 {
			continue
		}

		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1

			for j := i - 1; j >= 0; j-- {
				if ratings[j] > ratings[j+1] && candies[j] == candies[j+1] {
					candies[j] = candies[j+1] + 1
				} else {
					break
				}
			}
		}
	}

	sum := 0
	for _, c := range candies {
		sum += c
	}

	fmt.Println("RATINGS", ratings)
	fmt.Println("CANDIES", candies)

	return sum
}
