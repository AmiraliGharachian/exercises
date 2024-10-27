package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func addToSet(set map[int]bool, numbers []string) {
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			set[num] = true
		}
	}
}

func main() {

	fmt.Print("Enter first line of numbers: ")
	var input1 string
	fmt.Scanln(&input1)

	fmt.Print("Enter second line of numbers: ")
	var input2 string
	fmt.Scanln(&input2)

	numbers1 := strings.Split(input1, " ")
	numbers2 := strings.Split(input2, " ")

	numberSet := make(map[int]bool)

	addToSet(numberSet, numbers1)
	addToSet(numberSet, numbers2)

	union := make([]int, 0, len(numberSet))
	for num := range numberSet {
		union = append(union, num)
	}

	sort.Ints(union)

	fmt.Println("Union of numbers:")
	for _, num := range union {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
