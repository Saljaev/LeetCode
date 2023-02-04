func isIsomorphic(s string, t string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				if t[i] != t[j] {
					return false
				}
			}
		}
	}
	for i := 0; i < len(t); i++ {
		for j := i + 1; j < len(t); j++ {
			if t[i] == t[j] {
				if s[i] != s[j] {
					return false
				}
			}
		}
	}
	return true
}