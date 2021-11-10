package Arraysintoslice

// ConcatenateTwoArraysIntoSlice This function takes two arrays of undetermined length
// and returns a slice of the two arrays concatenated
func ConcatenateTwoArraysIntoSlice(array ...[]int) []int {
	var result []int

	for _, v := range array {
		result = append(result, v...)
	}

	return result
}
