func trap(height []int) int {
	waterArea := 0
	for i := 0; i < len(height)-1; i++ {
		area := 0
		if height[i] != 0 {
			size := height[i]
			for size > height[i+1] {
				for j := i + 1; j < len(height); j++ {
					if height[j] >= size && j-i > 1 {
						area = j - i - 1
						break
					}

				}
				size--
				waterArea += area
			}

		}

	}
	return waterArea
}