package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func excludeWords(words []string, excludeSet map[string]bool) []string {
	result := []string{}
	for _, word := range words {

		if !excludeSet[word] {
			result = append(result, word)
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the first line of words: ")
	scanner.Scan()
	firstLine := scanner.Text()
	words1 := strings.Fields(firstLine)

	fmt.Print("Enter the second line of words: ")
	scanner.Scan()
	secondLine := scanner.Text()
	words2 := strings.Fields(secondLine)

	excludeSet := make(map[string]bool)
	for _, word := range words2 {
		excludeSet[word] = true
	}

	result := excludeWords(words1, excludeSet)

	fmt.Println("Result:")
	fmt.Println(strings.Join(result, " "))
}
