package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 1, 1}
	sumOfUnique(arr)
}
func sumOfUnique(nums []int) int {
	store := make(map[int]int)
	for _, value := range nums {
		store[value]++
	}
	var sum int
	for num, count := range store {
		if count > 1 {
			sum += num * count
		}

	}
	fmt.Println(sum)
	return sum
}
