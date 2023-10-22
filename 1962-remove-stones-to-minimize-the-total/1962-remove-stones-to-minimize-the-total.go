func minStoneSum(piles []int, k int) int {
	stones := make(map[int]int)
	maxPile := 0
	minSum := 0
	for _, i := range piles {
		if i > maxPile {
			maxPile = i
		}
		stones[i]++
		minSum += i
	}
	i := 0
	for maxPile > 0 && i < k {
		if stones[maxPile] >= 1 {
			stones[maxPile]--
			minSum -= maxPile / 2
			stones[maxPile/2+maxPile%2]++
			i++
		} else {
			maxPile--
		}
	}

	return minSum
}