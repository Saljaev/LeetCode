func maxProfit(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			maxProfit += prices[i+1] - prices[i]
		}
	}
	return maxProfit
}
