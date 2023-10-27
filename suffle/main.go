package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Seed the random number generator.
	//	rand.Seed(time.Now().Unix())

	// Sample array (you can replace this with your own data).
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Shuffle the array.
	shuffleArray(array)

	// Print the shuffled array.
	fmt.Println(array)
}

func shuffleArray(arr []int) {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		// Generate a random index between 0 and i (inclusive).
		j := rand.Intn(i + 1)

		// Swap elements at indices i and j.
		arr[i], arr[j] = arr[j], arr[i]
	}
}
