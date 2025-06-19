package besttimetobuyandsellstockiv

func run(prices []int, k, currentProfit, idx, numTrx, prevPrice int, isBuying bool) int {

	if idx >= len(prices) {
		return currentProfit
	}

	if numTrx >= 2 {
		return currentProfit
	}

	if isBuying {
		return max(
			run(prices, k, currentProfit+prices[idx]-prevPrice, idx+1, numTrx+1, 0, false), // selling
			run(prices, k, currentProfit, idx+1, numTrx, prevPrice, isBuying),              // keep
		)
	} else {
		return max(
			run(prices, k, currentProfit, idx+1, numTrx, prices[idx], true),   // buying
			run(prices, k, currentProfit, idx+1, numTrx, prevPrice, isBuying), // keep
		)
	}
}

func maxProfit(k int, prices []int) int {
	return run(prices, k, 0, 0, 0, 0, false)
}
