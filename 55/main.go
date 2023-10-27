package main

import "fmt"

func canJump(nums []int) bool {
	maxReach := 0 // This variable stores the maximum index you can reach.

	for i := 0; i < len(nums); i++ {
		// If the current index is greater than the maximum reachable index, return false.
		if i > maxReach {
			return false
		}

		// Update maxReach to be the maximum of the current maxReach and the index we can jump to from the current position.
		maxReach = max(maxReach, i+nums[i])

		// If maxReach is greater than or equal to the last index, we can reach the end.
		if maxReach >= len(nums)-1 {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{2, 3, 1, 1, 4}
	fmt.Println("Example 1:", canJump(nums1)) // Output: true

	nums2 := []int{3, 2, 1, 0, 4}
	fmt.Println("Example 2:", canJump(nums2)) // Output: false

	nums3 := []int{2, 0, 0, 0, 4}
	fmt.Println("Example 3:", canJump(nums3)) // Output: false

	nums4 := []int{1, 1, 1, 0, 3, 2, 1, 0, 1}
	fmt.Println("Example 4:", canJump(nums4)) // Output: true

	nums5 := []int{1, 0, 1, 0, 1, 0, 1}
	fmt.Println("Example 5:", canJump(nums5)) // Output: false
}
