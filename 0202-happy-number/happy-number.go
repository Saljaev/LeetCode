package main

import (
	"fmt"
)

func isHappy(n int) bool {
	cycle := make(map[int]bool)
	start := n
	result := n
	for !cycle[start] {
		cycle[start] = true
		result = 0
		for start > 0 {
			result += (start % 10) * (start % 10)
			start /= 10
		}
		start = result
		if result == 1 {
			return true
		}
	}
	return false
}

func main() {
	test := 19
	fmt.Println(isHappy(test))
	fmt.Println(isHappy(1))
	fmt.Println(isHappy(7))
	fmt.Println(isHappy(3))
}
