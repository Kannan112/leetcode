package main

import "fmt"

func RemoveTwoNumbersAfterAPrime(array []int) []int {
	primeIndices := []int{} // Store indices of prime numbers

	// Find the indices of prime numbers
	for i, num := range array {
		if isPrime(num) {
			primeIndices = append(primeIndices, i)
		}
	}

	// Iterate through the prime indices and remove two numbers after each prime
	for _, idx := range primeIndices {
		if idx+2 < len(array) {
			array = append(array[:idx+1], array[idx+3:]...)
		}
	}

	return array
}

func main() {
	arr := []int{4, 2, 7, 3, 2, 8, 8, 0, 2}
	result := RemoveTwoNumbersAfterAPrime(arr)
	fmt.Println(result)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
