package factorialtrailingzeroes

func trailingZeroes(n int) int {

	fives := 0
	twos := 0

	for i := 1; i <= n; i++ {

		val := i

		for val%2 == 0 {
			val /= 2
			twos++
		}

		for val%5 == 0 {
			val /= 5
			fives++
		}

	}

	if fives < twos {
		return fives
	} else {
		return twos
	}
}
