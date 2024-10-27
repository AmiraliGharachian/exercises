package main

import "fmt"

func printGameInstructions() {
	fmt.Println("Welcome to Hangman!")
	fmt.Println("Instructions:")
	fmt.Println("1. Have the player select a letter of the alphabet.")
	fmt.Println("2. If the letter is contained in the word/phrase, all the letters equal to it are revealed.")
	fmt.Println("3. If the letter is not contained in the word/phrase, a portion of the hangman is added.")
	fmt.Println("4. The game continues until:")
	fmt.Println("   1. The word/phrase is guessed (all letters are revealed) - WINNER")
	fmt.Println("   2. All the parts of the hangman are displayed - LOSER")
}

func main() {

	printGameInstructions()
}
