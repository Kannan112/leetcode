package main

import "fmt"

func main() {
	nums1 := []int{5, 4, 7, 2}
	nums2 := []int{0, 6, 1, 2}
	output := intersect(nums1, nums2)
	fmt.Println(output)
}

func intersect(nums1 []int, nums2 []int) []int {
	// Create a map to store the count of each number in nums1
	countMap := make(map[int]int)

	// Populate the countMap with numbers from nums1
	for _, num := range nums1 {
		countMap[num]++
	}

	// Initialize a result slice to store the intersection
	result := make([]int, 0)

	// Iterate through nums2 and check for intersections
	for _, num := range nums2 {
		if count, found := countMap[num]; found && count > 0 {
			result = append(result, num)
			countMap[num]--
		}
	}

	return result
}
