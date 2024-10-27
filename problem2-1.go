package main

import (
	"fmt"
	"math"
)

func main() {
	var N int

	fmt.Print("Enter the count of numbers (N): ")
	fmt.Scan(&N)

	numbers := make([]int, N)

	fmt.Println("Enter the numbers:")
	for i := 0; i < N; i++ {
		fmt.Scan(&numbers[i])
	}

	for i := 0; i < N; i++ {
		numbers[i] = int(math.Pow(float64(numbers[i]), 2))
	}

	fmt.Println("Numbers raised to the power of 2:")
	for _, num := range numbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
