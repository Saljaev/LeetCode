func numberOfCombination(k, n int) int {
	if k == 0 || n == 0 {
		return 1
	} else {
		return numberOfCombination(k-1, n-1) * n / k
	}
}

func climbStairs(n int) int {
	ways := 1
	i := 1
	for i*2 <= n {
		ways += numberOfCombination(i, n-i*2+i)
		i++
		
	}
	return ways
}