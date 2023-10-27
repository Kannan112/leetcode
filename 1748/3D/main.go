// package main

// import "fmt"

// func main() {
// 	// Define the dimensions of the 3D array
// 	xSize := 3
// 	ySize := 4
// 	zSize := 2

// 	// Create a 3D array using nested slices
// 	var array3D [][][]int

// 	// Initialize the 3D array with values
// 	for x := 0; x < xSize; x++ {
// 		// Create a new 2D slice for each x
// 		var matrix2D [][]int
// 		for y := 0; y < ySize; y++ {
// 			// Create a new 1D slice for each y
// 			var row []int
// 			for z := 0; z < zSize; z++ {
// 				// Append values to the row
// 				row = append(row, x*100+y*10+z)
// 			}
// 			// Append the row to the 2D matrix
// 			matrix2D = append(matrix2D, row)
// 		}
// 		// Append the 2D matrix to the 3D array
// 		array3D = append(array3D, matrix2D)
// 	}

//		// Access and print values from the 3D array
//		for x := 0; x < xSize; x++ {
//			for y := 0; y < ySize; y++ {
//				for z := 0; z < zSize; z++ {
//					fmt.Printf("array3D[%d][%d][%d] = %d\n", x, y, z, array3D[x][y][z])
//				}
//			}
//		}
//	}
package main

func main() {
	var array3D [][][]int
}
