package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}

	out := make([][]int, len(matrix[0]))
	for i := range out {
		out[i] = make([]int, len(matrix))
	}

	for i, v := range matrix {
		for j, v1 := range v {
			out[j][i] = v1
		}
	}
	return out
}

func main() {
	test := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	test1 := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(transpose(test))
	fmt.Println(transpose(test1))
}
