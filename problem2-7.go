package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Enter some text: ")
	fmt.Scanln(&input)

	words := strings.Fields(input)

	longestWord := ""
	longestLength := 0

	for _, word := range words {
		wordLength := len(word)

		if wordLength > longestLength {
			longestLength = wordLength
			longestWord = word
		}
	}

	fmt.Println(longestWord)
}
