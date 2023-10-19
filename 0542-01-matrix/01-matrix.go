package main

import "fmt"

func updateMatrix(mat [][]int) [][]int {
	result := mat
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				result[i][j] = 0
			} else {
				if (i-1 > 0 && mat[i-1][j] == 0) || (i+1 < len(mat) && mat[i+1][j] == 0) || (j-1 > 0 && mat[i][j-1] == 0) || (j+1 < len(mat[i]) && mat[i][j+1] == 0) {
					result[i][j] = 1
				} else {
					if i-1 > 0 {
						result[i][j] = result[i-1][j] + 1
					} else if j+1 < len(mat[i]) {
						result[i][j] = result[i][j+1] + 1
					}
				}
			}
		}
	}
	return result
}

func main() {
	test := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	fmt.Println(updateMatrix(test))
}
