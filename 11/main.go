package main

import "fmt"

func main() {
	array := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	quantity := maxArea(array)
	fmt.Println(quantity)
}
func maxArea(height []int) []int {
	quantity := []int{}
	LeftBrick := height[0]
	RightBrick := height[len(height)]
	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j > 0; j-- {
			if LeftBrick <= RightBrick {
				quantity = append(quantity, LeftBrick*j)

			}
		}
	}
	return quantity
}
