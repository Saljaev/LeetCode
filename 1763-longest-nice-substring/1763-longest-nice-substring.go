func longestNiceSubstring(s string) string {
	longestNice := ""

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if isNice(s[i:j+1]) && j+1-i > len(longestNice) {
				longestNice = s[i : j+1]
			}
		}
	}

	return longestNice
}

func isNice(s string) bool {
	if len(s) < 2 {
		return false
	}

	uppercaseSet := make(map[rune]bool)
	lowercaseSet := make(map[rune]bool)

	for _, char := range s {
		if unicode.IsLower(char) {
			lowercaseSet[char] = true
		} else {
			uppercaseSet[char] = true
		}
	}

	for k, _ := range lowercaseSet {
		if !uppercaseSet[unicode.ToUpper(k)] {
			return false
		}
	}

	for k, _ := range uppercaseSet {
		if !lowercaseSet[unicode.ToLower(k)] {
			return false
		}
	}

	return true
}