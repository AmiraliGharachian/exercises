package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var N, X int

	fmt.Print("Enter the range (N): ")
	fmt.Scanln(&N)

	if N < 10 {
		fmt.Println("Error: The range (N) must be at least 10.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(N-10+1) + 10

	fmt.Print("Guess the number (X): ")
	fmt.Scanln(&X)

	if X == randomNumber {
		fmt.Println("You guessed the number correctly!")
	} else {
		fmt.Printf("That's not the number I had in mind. You guessed %d but my number was %d.\n", X, randomNumber)
	}
}
