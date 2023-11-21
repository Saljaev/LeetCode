func permute(nums []int) [][]int {
	var result [][]int
	backtrack(&result, []int{}, nums)
	return result
}

func backtrack(result *[][]int, tempList []int, nums []int) {
	if len(tempList) == len(nums) {
		temp := make([]int, len(tempList))
		copy(temp, tempList)
		*result = append(*result, temp)
		return
	} else {
		for i := 0; i < len(nums); i++ {
			if contains(tempList, nums[i]) {
				continue
			}
			tempList = append(tempList, nums[i])
			backtrack(result, tempList, nums)
			tempList = tempList[:len(tempList)-1]
		}
	}
}

func contains(arr []int, elem int) bool {
	for _, value := range arr {
		if value == elem {
			return true
		}
	}
	return false
}