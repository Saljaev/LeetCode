func sortArray(nums []int) []int {
	quickSort(nums)
	return nums
}

func quickSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	v := nums[0]
	less, bigger := 0, 0
	for i := 1; i < len(nums)-bigger; {
		if nums[i] < v {
			nums[less], nums[i] = nums[i], nums[less]
			less++
		} else if nums[i] > v {
			nums[i], nums[len(nums)-bigger-1] = nums[len(nums)-bigger-1], nums[i]
			bigger++
		} else {
			i++
		}
	}

	quickSort(nums[:less])
	quickSort(nums[len(nums)-bigger:])
}
