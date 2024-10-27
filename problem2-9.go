package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var target int

	fmt.Print("Enter the numbers: ")
	fmt.Scanln(&input)

	fmt.Print("Enter the target sum N: ")
	fmt.Scan(&target)

	nums := strings.Fields(input)
	intNums := make([]int, len(nums))

	for i, num := range nums {
		var n int
		fmt.Sscan(num, &n)
		intNums[i] = n
	}

	found := false
	for i := 0; i < len(intNums); i++ {
		sum := 0
		subset := []string{}
		for j := i; j < len(intNums); j++ {
			sum += intNums[j]
			subset = append(subset, fmt.Sprintf("%d", intNums[j]))
			if sum == target {
				fmt.Println(strings.Join(subset, "+"))
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		fmt.Println("no solution exists")
	}
}
