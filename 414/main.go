package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{3, 2}
	number := thirdMax(array)
	fmt.Println(number)
}
func thirdMax(nums []int) int {
	firstMax, secondMax, thirdMax := math.MinInt64, math.MinInt64, math.MinInt64

	for _, num := range nums {
		if num > firstMax {
			thirdMax = secondMax
			secondMax = firstMax
			firstMax = num
		} else if num > secondMax && num < firstMax {
			thirdMax = secondMax
			secondMax = num
		} else if num > thirdMax && num < secondMax {
			thirdMax = num
		}
	}

	if thirdMax != math.MinInt64 {
		return thirdMax
	} else {
		return firstMax
	}
}
