func runningSum(nums []int) []int {
	resultArr := []int{}
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			resultArr = append(resultArr, nums[i]+resultArr[i-1])
		} else {
			resultArr = append(resultArr, nums[i])
		}
	}
	return resultArr
}