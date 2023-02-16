func maxProfit(prices []int) int {
	profitOne := 0
	profit := 0
	profitTwo := 0
	min := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if profitOne < prices[i]-min {
			profitOne = prices[i] - min

		}
	}
	min = prices[0]
	for i := 1; i < len(prices)-1; i++ {
		if prices[i] < min {
			min = prices[i]

			// fmt.Println(min)
		} else if prices[i+1] < prices[i] {
			profitTwo += prices[i] - min
			min = prices[i+1]

			// fmt.Println(min)
			// fmt.Println(prices[i], min)
		} else if i == len(prices)-2 && min < prices[i+1] {
			profitTwo += prices[i+1] - min
		}
	}
	if profitOne > profitTwo {
		profit = profitOne
	} else {
		profit = profitTwo
	}
	return profit
}