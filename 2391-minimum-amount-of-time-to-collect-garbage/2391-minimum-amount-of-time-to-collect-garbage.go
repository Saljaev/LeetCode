func garbageCollection(garbage []string, travel []int) int {
	result := 0
	posMetal, posPaper, posGlass := 0, 0, 0
	for j := 0; j < len(garbage); j++ {
		for k := 0; k < len(garbage[j]); k++ {
			if garbage[j][k] == 'M' {
				posMetal = j
			} else if garbage[j][k] == 'P' {
				posPaper = j
			} else if garbage[j][k] == 'G' {
				posGlass = j
			}
			result++
		}
	}
	for i := 0; i < len(travel); i++ {
		if i < posMetal {
			result += travel[i]
		}
		if i < posGlass {
			result += travel[i]
		}
		if i < posPaper {
			result += travel[i]
		}
	}
	return result
}