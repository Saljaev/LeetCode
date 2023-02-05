func longestPalindrome(s string) int {
    frequency := map[string]int{}
	for i := 0; i < len(s); i++ {
		frequency[string(s[i])] += 1
	}
	lenght := 0
	odd := false
	for _, value := range frequency {
		lenght += value / 2 * 2
		if value%2 == 1 {
			odd = true
		}
	}
	if odd {
		lenght++
	}
	return lenght
}