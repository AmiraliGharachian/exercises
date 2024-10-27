package main

import (
	"fmt"
)

type Point struct {
	row int
	col int
}

func PrintAt(p Point, message string) {

	fmt.Printf("Printing at Row: %d, Column: %d -> %s\n", p.row, p.col, message)
}

func main() {

	point1 := Point{row: 5, col: 5}

	PrintAt(point1, "Hello, World!")

	point2 := Point{row: 10, col: 20}
	PrintAt(point2, "Welcome to Go!")
}
