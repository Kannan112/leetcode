package main

import "fmt"

func main() {
	array := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//snew := [][]string{}

	doublearray := [][]string{{"eat", "tea"}, {"tan", "ate", "nat"}, {"bat"}}
	//fmt.Println(groupAnagrams(array))
	//Array(doublearray, array)
	for first, _ := range doublearray {
		fmt.Println(first)
	}
	for i, j := range array {
		sk := make([]int, len(j))
	}
}

// func groupAnagrams(strs []string) [][]string {
// 	var result [][]string
// 	for i := 0; i < len(strs); i++ {
// 		for j := 0; j < len(strs); j++ {
// 			letter := strs[i]

// 		}
// 	}
// 	return result
// }

// func Array(doub [][]string, array []string) {
// 	fmt.Println(len(doub))
// 	fmt.Println(array[1])
// 	for i := 0; i < len(doub); i++ {
// 		fmt.Println(doub[i])
// 	}
// 	for _, str := range array {
// 		asciiVals := make([]int, len(str))
// 		for i, char := range str {
// 			asciiVals[i] = int(char)
// 		}
// 		fmt.Printf("ASCII values of %s: %v\n", str, asciiVals)
// 	}

// }
