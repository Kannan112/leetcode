package main

import (
	"fmt"
)

func decimalToBinary(number int) string {
	binary = fmt.Sprintf("%d%s", remainder, binary)
	return binary
}

func main() {
	var num int
	fmt.Print("Enter a decimal number: ")
	fmt.Scan(&num)

	binary := decimalToBinary(num)
	fmt.Printf("Binary representation of %d is %s\n", num, binary)
}
