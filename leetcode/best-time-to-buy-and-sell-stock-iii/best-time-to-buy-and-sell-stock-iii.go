package besttimetobuyandsellstockiii

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProfit2(prices []int) int {

	N := len(prices)

	maxPro := 0

	buy1 := prices[0]
	for x := range N {
		trx1 := prices[x] - buy1
		buy1 = min(buy1, prices[x])
		maxPro = max(maxPro, trx1)

		buy2 := 999999

		for y := x + 1; y < N; y++ {

			trx2 := trx1 + prices[y] - buy2
			maxPro = max(maxPro, trx2)

			if y > x {
				buy2 = min(buy2, prices[y])
			}
		}
	}

	return maxPro
}

func maxProfit(prices []int) int {

	N := len(prices)

	maxPro := 0

	buy1 := prices[0]
	for x := range N {
		trx1 := prices[x] - buy1
		buy1 = min(buy1, prices[x])
		maxPro = max(maxPro, trx1)

		buy2 := 999999
		for y := x + 1; y < N; y++ {

			trx2 := trx1 + prices[y] - buy2
			maxPro = max(maxPro, trx2)

			if y > x {
				buy2 = min(buy2, prices[y])
			}
		}
	}

	return maxPro
}
