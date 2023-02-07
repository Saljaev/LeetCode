func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		average := (right + left) / 2
		if nums[average] > target {
			right = average - 1
		} else if nums[average] < target {
			left = average + 1
		} else {
			return average
		}

	}
	return -1
}