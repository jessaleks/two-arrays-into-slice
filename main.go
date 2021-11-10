package main

import "fmt"

// conctenateTwoArraysIntoSlice
// This function takes two arrays of undertermined length and returns a slice of the two arrays concatenated
func concatenateTwoArraysIntoSlice(array ...[]int) []int {
	var result []int

	for _, v := range array {
		result = append(result, v...)
	}

	return result
}

func main() {
	intArray1 := [5]int{1, 2, 3, 4, 5}
	intArray2 := [5]int{1, 2, 3, 4, 5}

	var intSlice []int

	intSlice = concatenateTwoArraysIntoSlice(intArray1[:], intArray2[:])
	fmt.Println(intSlice)
}
