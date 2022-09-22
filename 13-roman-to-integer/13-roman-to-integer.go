func romanToInt(s string) int {
	countI := 0
	countX := 0
	countC := 0
	result := 0
	for i := 0; i < utf8.RuneCountInString(s); i++ {
		switch s[i] {
		case 'I':
			countI++
			result++
		case 'V':
			result += 5 - 2*countI
			countI = 0
		case 'X':
			result += 10 - 2*countI
			countI = 0
			countX += 10
		case 'L':
			result += 50 - 2*countX
			countX = 0
		case 'C':
			result += 100 - 2*countX
			countX = 0
			countC += 100
		case 'D':
			result += 500 - 2*countC
			countC = 0
		case 'M':
			result += 1000 - 2*countC
			countC = 0
		}

	}
	return result
}