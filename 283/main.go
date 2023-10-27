package main

func main() {
	arr := []int{0, 2, 1, 4, 0, 1}
	moveZeroes(arr)
}

func moveZeroes(nums []int) {
	nonZeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// Swap non-zero element with the first zero element (if any)
			nums[i], nums[nonZeroIndex] = nums[nonZeroIndex], nums[i]
			nonZeroIndex++
		}
	}
}
