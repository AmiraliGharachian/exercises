package main

import (
	"log"

	"github.com/bugs/asterio"
)

func PrintAt(row int, col int, str string) {

	asterio.MoveCursor(row, col)

	_, err := asterio.Print(str)
	if err != nil {
		log.Fatalf("Failed to print string: %v", err)
	}
}

func main() {

	PrintAt(5, 5, "hello")
}
