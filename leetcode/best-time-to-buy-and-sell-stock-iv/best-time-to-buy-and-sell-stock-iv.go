package besttimetobuyandsellstockiv

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

var memMinPrices [][]*int
var memMaxProfit [][]*int

func minPrice(prices []int, a int, b int) int {

	if memMinPrices[a][b] != nil {
		return *memMinPrices[a][b]
	}

	res := 0
	if a == b {
		res = prices[a]
	} else {
		res = min(prices[a], minPrice(prices, a+1, b))
	}
	memMinPrices[a][b] = &res
	return res
}

func maxProfitKN(prices []int, k, n int) int {
	if k <= 0 {
		return 0
	}

	if n <= 0 {
		return 0
	}

	if memMaxProfit[k][n] != nil {
		return *memMaxProfit[k][n]
	}

	maxPro := maxProfitKN(prices, k, n-1)

	for j := range k {
		maxPro = max(maxPro, maxProfitKN(prices, j, n))
	}

	for m := range n {
		lastProfit := prices[n] - minPrice(prices, m, n-1)
		tmpMaxPro := maxProfitKN(prices, k-1, m-1) + lastProfit
		maxPro = max(maxPro, tmpMaxPro)
	}

	memMaxProfit[k][n] = &maxPro

	return maxPro
}

func maxProfit(k int, prices []int) int {
	n := len(prices)

	memMinPrices = make([][]*int, n)
	for i := range n {
		memMinPrices[i] = make([]*int, n)
	}

	memMaxProfit = make([][]*int, k+1)
	for i := range k + 1 {
		memMaxProfit[i] = make([]*int, n)
	}

	return maxProfitKN(prices, k, n-1)
}
