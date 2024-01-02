package main

import "fmt"

func unequalTriplets(nums []int) int {
	triplets := 0
	pairs := 0
	count := [1001]int{}
	for i := 0; i < len(nums); i++ {
		triplets += pairs - count[nums[i]]*(i-count[nums[i]])
		pairs += i - count[nums[i]]
		count[nums[i]]++
	}
	return triplets
}

func main() {
	numbers := []int{4, 4, 2, 4, 3}
	unique := unequalTriplets(numbers)
	fmt.Println(unique)
}
