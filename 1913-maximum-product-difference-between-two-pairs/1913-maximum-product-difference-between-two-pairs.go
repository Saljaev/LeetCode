func maxProductDifference(nums []int) int {
	maxFirst, maxSecond, minFirst, minSecond := nums[0], nums[1], nums[0], nums[1]
	for _, num := range nums[1:] {
		if num > maxFirst {
			maxSecond = maxFirst
			maxFirst = num
		} else if num > maxSecond {
			maxSecond = num
		}

		if num < minFirst {
			minSecond = minFirst
			minFirst = num
		} else if num < minSecond {
			minSecond = num
		}
	}
	return (maxFirst * maxSecond) - (minFirst * minSecond)
}