func isSubsequence(s string, t string) bool {
	arr := []int{}
	lastpos := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				if j > lastpos || lastpos == 0 {
					arr = append(arr, j)
					lastpos = j
					break
				}
			}
		}
		if len(arr) == i {
			return false
		}
	}
	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}