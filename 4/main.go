// package main

// import "fmt"

//	func main() {
//		arr1 := []int{}
//		arr2 := []int{2, 3}
//		median := findMedianSortedArrays(arr1, arr2)
//		fmt.Println(median)
//	}
//
//	func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//		var merge []int
//		merge = append(merge, nums1...)
//		merge = append(merge, nums2...)
//		for i := 0; i < len(merge); i++ {
//			for j := 0; j < len(merge); j++ {
//				if merge[i] > merge[j] {
//					merge[i], merge[j] = merge[j], merge[i]
//				}
//			}
//		}
//		if len(merge) == 1 {
//			return float64(merge[0])
//		}
//		center := len(merge) / 2
//		if center%2 == 0 {
//			EvenResult := float64(merge[center]+merge[center-1]) / 2
//			fmt.Println(EvenResult)
//			return EvenResult
//		}
//		return float64(merge[center])
//	}
package main

import (
	"fmt"
)

func main() {
	arr1 := []int{}
	arr2 := []int{2, 3}
	median := findMedianSortedArrays(arr1, arr2)
	fmt.Println(median)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Merge the two sorted arrays
	merge := mergeSortedArrays(nums1, nums2)

	// Find the median
	length := len(merge)
	if length%2 == 0 {
		// Even number of elements, return the average of the two middle elements
		center := length / 2
		median := float64(merge[center-1]+merge[center]) / 2.0
		return median
	} else {
		// Odd number of elements, return the middle element
		center := length / 2
		median := float64(merge[center])
		return median
	}
}

func mergeSortedArrays(nums1, nums2 []int) []int {
	merged := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0

	// Merge the two sorted arrays into a single sorted array
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	// Append any remaining elements from both arrays
	merged = append(merged, nums1[i:]...)
	merged = append(merged, nums2[j:]...)

	return merged
}
