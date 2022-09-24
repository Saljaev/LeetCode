func isPalindrome(x int) bool {
	x1 := x
	ar := []int{}
	if x < 0 {
		return false
	}
	for x1 > 0 {
		ar = append(ar, x1%10)
		x1 /= 10
	}
	for i := len(ar) - 1; i > 0; i-- {
		if ar[i] != x%10 {
			return false
		}
		x /= 10
	}
	return true
}