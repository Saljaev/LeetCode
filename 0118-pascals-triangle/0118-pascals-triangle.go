func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		arr := make([]int, i+1)
		arr[0] = 1
		arr[i] = 1
		for j := 1; j < len(result[i-1]); j++ {
			arr[j] = result[i-1][j-1] + result[i-1][j]
		}
		result[i] = arr
	}

	return result
}