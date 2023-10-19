package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	position := 0
	for position < len(nums) {
		if nums[position] == target || nums[position] > target {
			return position
		}
		if position >= 1 {
			if target > nums[position-1] && target < nums[position] {
				return position
			}
		}
		position++
	}
	return len(nums)
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Println(searchInsert([]int{1}, 2))
	fmt.Println(searchInsert([]int{1, 3}, 1))
}
