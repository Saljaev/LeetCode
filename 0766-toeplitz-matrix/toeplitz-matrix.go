package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}

func main() {
	test := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	test1 := [][]int{{1, 2}, {2, 2}}
	fmt.Println(isToeplitzMatrix(test))
	fmt.Println(isToeplitzMatrix(test1))
}
