func pivotIndex(nums []int) int {
	index := 0
	for i := 0; i <= len(nums); i++ {
		if i > 0 {
			if sum(nums[:i-1]) == sum(nums[i:]) {
				return index - 1
			}
		}
		index++
	}
	return -1
}

func sum(arr []int) int {
	sumarr := 0
	for i := 0; i < len(arr); i++ {
		sumarr += arr[i]
	}
	return sumarr
}