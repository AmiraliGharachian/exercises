package main

import (
	"fmt"
	"strconv"
	"strings"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {

		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {

	fmt.Print("Enter numbers separated by space: ")
	var input string
	fmt.Scanln(&input)

	inputNumbers := strings.Split(input, " ")
	arr := make([]int, len(inputNumbers))

	for i, numStr := range inputNumbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Invalid number:", numStr)
			return
		}
		arr[i] = num
	}

	selectionSort(arr)

	fmt.Println("Sorted array:")
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
