func maximumNumberOfStringPairs(words []string) int {
	result := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if words[i][1] == words[j][0] && words[i][0] == words[j][1] {
				result++
				break
			}
		}

	}
	return result
}