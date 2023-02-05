func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]
	for i := 0; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else {
			if profit < prices[i]-minPrice {
				profit = prices[i] - minPrice
			}
		}
	}
	return profit
}