package main

// import "fmt"

// func main() {
// 	nums := []int{41900, 69441, 94407, 37498, 20299, 10856, 36221, 2231, 54526, 79072, 84309, 76765, 92282, 13401, 44698, 17586, 98455, 47895, 98889, 65298, 32271, 23801, 83153, 12186, 7453, 79460, 67209, 54576, 87785, 47738, 40750, 31265, 77990, 93502, 50364, 75098, 11712, 80013, 24193, 35209, 56300, 85735, 3590, 24858, 6780, 50086, 87549, 7413, 90444, 12284, 44970, 39274, 81201, 43353, 75808, 14508, 17389, 10313, 90055, 43102, 18659, 20802, 70315, 48843, 12273, 78876, 36638, 17051, 20478}
// 	fmt.Println(minimumDifference(nums, 5))
// }

// func minimumDifference(nums []int, k int) int {
// 	if k == 0 {
// 		return 0
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums); j++ {
// 			if nums[i] > nums[j] {
// 				nums[i], nums[j] = nums[j], nums[i]
// 			}
// 		}
// 	}
// 	index := k - 1
// 	result := nums[0] - nums[index]
// 	return result
// }
import (
	"fmt"
	"sort"
)

func minDifference(nums []int, k int) int {
	if k >= len(nums) {
		return 0
	}

	sort.Ints(nums)
	minDiff := int(^uint(0) >> 1) // Initialize minDiff to maximum int value

	for i := 0; i <= len(nums)-k; i++ {
		minDiff = min(minDiff, nums[i+k-1]-nums[i])
	}

	return minDiff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Test cases
func main() {
	nums := []int{41900, 69441, 94407, 37498, 20299, 10856, 36221, 2231, 54526, 79072, 84309, 76765, 92282, 13401, 44698, 17586, 98455, 47895, 98889, 65298, 32271, 23801, 83153, 12186, 7453, 79460, 67209, 54576, 87785, 47738, 40750, 31265, 77990, 93502, 50364, 75098, 11712, 80013, 24193, 35209, 56300, 85735, 3590, 24858, 6780, 50086, 87549, 7413, 90444, 12284, 44970, 39274, 81201, 43353, 75808, 14508, 17389, 10313, 90055, 43102, 18659, 20802, 70315, 48843, 12273, 78876, 36638, 17051, 20478}

	k1 := 5
	fmt.Println(minDifference(nums, k1)) // Output: 0

	nums2 := []int{9, 4, 1, 7}
	k2 := 2
	fmt.Println(minDifference(nums2, k2)) // Output: 2
}
