package main

import "fmt"

func main() {
	value := "NsESW"

	check := isPathCrossing(value)
	fmt.Println(check)
}

func isPathCrossing(path string) bool {
	result := make(map[[2]int]bool)
	result[[2]int{0, 0}] = true
	y, x := 0, 0
	for _, char := range path {
		if char == 'N' {
			y++
		} else if char == 'S' {
			y--
		} else if char == 'W' {
			x++
		} else if char == 'E' {
			x--
		}
		if result[[2]int{x, y}] == true {
			return true
		}
		result[[2]int{x, y}] = true

	}
	return false

}
