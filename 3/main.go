package main

import (
	"fmt"
)

func main() {
	name := "abcdababcde"
	result := lengthOfLongestSubstring(name)
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	var Array []int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			store := (s[i])
			Array := append(Array, store)
		}
	}
}
