func countSubarrays(nums []int, minK, maxK int) int64 {
	maxPos, minPos, bound := -1, -1, -1
	result := 0
	for i, item := range nums {
		if item < minK || item > maxK {
			bound = i
		}
		if item == maxK {
			maxPos = i
		}
		if item == minK {
			minPos = i
		}
		result += max(0, min(minPos, maxPos)-bound)
	}
	return int64(result)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}