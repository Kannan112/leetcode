package main

import "fmt"

func main() {
	array := []int{1, 2, 1, 5, 2}
	result := singleNumber(array)
	fmt.Println(result)
}

func singleNumber(nums []int) int {
	maping := make(map[int]int)
	for _, value := range nums {
		maping[value]++
	}
	for value, count := range maping {
		if count <= 1 {
			return value
		}
	}
	return 0
}
