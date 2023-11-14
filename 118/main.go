package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0], result[i][i] = 1, 1

		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	return result
}

func main() {
	numRows := 5 // You can change this to any desired number of rows
	pascalsTriangle := generate(numRows)

	// Print the Pascal's Triangle
	for _, row := range pascalsTriangle {
		fmt.Println(row)
	}
}
