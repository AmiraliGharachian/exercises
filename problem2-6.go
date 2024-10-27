package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Enter some text: ")
	fmt.Scanln(&input)

	input = strings.ToLower(input)

	letterCount := make(map[rune]int)

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			letterCount[char]++
		}
	}

	for letter := 'a'; letter <= 'z'; letter++ {
		fmt.Printf("%c: %d\n", letter, letterCount[letter])
	}
}
