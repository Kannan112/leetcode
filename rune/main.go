package main

import "fmt"

func areOccurrencesEqual(s string) bool {
	frequency := make(map[rune]int)
	for o, char := range s {
		frequency[char]++
		fmt.Println("char", char, "o", o)
	}
	firstCharFrequency := frequency[rune(s[0])]

	for _, freq := range frequency {
		fmt.Println(firstCharFrequency, freq)
		if freq != firstCharFrequency {
			return false
		}
	}
	return true

}

func main() {
	s := "aaabbb"
	fmt.Println(areOccurrencesEqual(s))
}
