package main

import (
	"fmt"
	"github.com/jessaleks/two-arrays-into-slice/arraysIntoSlice"
)


func main() {
	intArray1 := [5]int{1, 2, 3, 4, 5}
	intArray2 := [5]int{1, 2, 3, 4, 5}

	var intSlice []int

	intSlice = Arraysintoslice.ConcatenateTwoArraysIntoSlice(intArray1[:], intArray2[:])
	fmt.Println(intSlice)
}
