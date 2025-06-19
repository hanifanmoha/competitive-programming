package besttimetobuyandsellstockiii

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func run(prices []int, currentProfit, idx, numTrx, prevPrice int, isBuying bool) int {

	if idx >= len(prices) {
		return currentProfit
	}

	if numTrx >= 2 {
		return currentProfit
	}

	if isBuying {
		return max(
			run(prices, currentProfit+prices[idx]-prevPrice, idx+1, numTrx+1, 0, false), // selling
			run(prices, currentProfit, idx+1, numTrx, prevPrice, isBuying),              // keep
		)
	} else {
		return max(
			run(prices, currentProfit, idx+1, numTrx, prices[idx], true),   // buying
			run(prices, currentProfit, idx+1, numTrx, prevPrice, isBuying), // keep
		)
	}
}

func maxProfit(prices []int) int {
	return run(prices, 0, 0, 0, 0, false)
}
