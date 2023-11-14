package main

import "fmt"

func main() {
	s := "hello world it me"
	a := truncateSentence(s, 2)
	fmt.Println(a)
}
func truncateSentence(s string, k int) string {
	limit := 0
	var result string
	for i := 0; i < len(s); i++ {
		if k == 0 {
			return " hello"
		}
		if k > limit {
			if k == limit-1 {
				return result
			}
			result += string(s[i])
			if s[i] == ' ' {
				limit++
			}
		}
	}
	return result
}
