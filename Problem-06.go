package main

import (
	"fmt"
	"strings"
)

func main() {
	var character string

	fmt.Print("Enter a character name: ")
	fmt.Scanln(&character)

	switch strings.ToLower(character) {
	case "geralt", "ciri", "yennefer":
		fmt.Printf("%s is part of the game The Witcher 3!\n", character)
	case "aldaris", "jim raynor", "kerrigan", "zeratul":
		fmt.Printf("%s is part of the game StarCraft 2!\n", character)
	case "arthas", "illidan", "sylvanas":
		fmt.Printf("%s is part of the game WarCraft 3!\n", character)
	default:
		fmt.Println("I don't recognize this character...")
	}
}
