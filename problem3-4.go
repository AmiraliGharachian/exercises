package main

import (
	"fmt"
)

func incrementBy(slice []int, n int) []int {
	for i := range slice {
		slice[i] += n
	}
	return slice
}

func main() {
	inputSlice := []int{1, 2, 3, 4}

	result := incrementBy(inputSlice, 5)
	fmt.Println(result)
}
