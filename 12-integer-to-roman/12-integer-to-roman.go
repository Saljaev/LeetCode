func intToRoman(num int) string {
	s := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	result := ""
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	index := 0
	minus := s[roman[index]]
	for num > 0 {
		if num-minus >= 0 {
			num -= minus
			result += roman[index]

		} else {
			index++
			minus = s[roman[index]]
		}

	}

	return result
}