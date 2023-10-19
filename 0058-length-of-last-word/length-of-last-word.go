package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	result := 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == ' ' && s[i+1] != ' ' {
			result = 1
		} else if s[i+1] != ' ' {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(lengthOfLastWord("Hello  world"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
	fmt.Println(lengthOfLastWord("f"))
	fmt.Println(lengthOfLastWord(" f"))
}
