func timeRequiredToBuy(tickets []int, k int) int {
	second := 0
	for tickets[k] > 0 {
		for j := 0; j < len(tickets); j++ {
			if tickets[j] > 0 {
				second++
				tickets[j]--
			}
			if tickets[k] == 0 {
				break
			}
		}
	}
	return second
}