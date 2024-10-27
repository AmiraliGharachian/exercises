package main

import (
	"fmt"
)

func subslice(slice []int, start int, end int) []int {
	if start < 0 || end > len(slice) || start >= end {
		return []int{}
	}
	return slice[start:end]
}

func main() {

	inputSlice := []int{1, 2, 3, 4, 5, 6, 7}

	result := subslice(inputSlice, 1, 4)
	fmt.Println(result)
}
