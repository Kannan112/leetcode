package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	merge(arr, 3, arr2, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	//var Output []int
	nums2 = nums2[:n]
	nums1 = nums1[:m]
	nums1 = append(nums1, nums2...)
	// one := nums1[:m]
	// two := nums2[:n]

	// Output = append(Output, one...)
	// Output = append(Output, two...)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1); j++ {
			if nums1[i] < nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}
	fmt.Println(nums1)
}
