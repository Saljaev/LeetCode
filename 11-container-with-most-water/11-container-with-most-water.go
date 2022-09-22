func maxArea(height []int) int {
	left := 0
	rigth := len(height) - 1
	maxArea := 0
	area := 0
	for left < rigth {
		if height[left] > height[rigth] {
			area = height[rigth] * (rigth - left)
			rigth--
		} else {
			area = height[left] * (rigth - left)
			left++
		}
		if area > maxArea {
			maxArea = area
		}

	}
	return maxArea
}