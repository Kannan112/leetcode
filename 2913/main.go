package main

import "fmt"

func main() {
	array := []int{1, 2, 1}
	result := sumCounts(array)
	fmt.Println(result)
}

func sumCounts(nums []int) int {
	set1 := len(nums) * len(nums)
	maping := make(map[int]int)
	for _, value := range nums {
		maping[value]++
	}
	for _, num := range maping {
		sum := num * num
	}
	set1 += sum
	return set1
}
