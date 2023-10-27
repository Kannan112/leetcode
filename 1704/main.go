package main

import (
	"fmt"
)

func main() {
	something := "helloherlos"
	check := halvesAreAlike(something)
	fmt.Println(check)
}
func halvesAreAlike(s string) bool {
	vowels := "aeiouAEIOU"
	//store := make(map[string]int)
	lengthOfString := len(s)
	half := lengthOfString / 2
	a := s[:half]
	b := s[half:]
	var ACount, BCount int
	for i := 0; i < half; i++ {
		for j := 0; j < len(vowels); j++ {
			if a[i] == vowels[j] {
				ACount++
			}
			if b[i] == vowels[j] {
				BCount++
			}
		}
	}
	if ACount != BCount {
		return false
	}
	return true
}
