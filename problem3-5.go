package main

import (
	"fmt"
)

type Point struct {
	row int
	col int
}

func main() {

	point1 := Point{row: 1, col: 2}
	point2 := Point{row: 3, col: 4}
	point3 := Point{row: 5, col: 6}

	fmt.Println("Point 1:", point1)
	fmt.Println("Point 2:", point2)
	fmt.Println("Point 3:", point3)

	fmt.Printf("Point 1 - Row: %d, Col: %d\n", point1.row, point1.col)
	fmt.Printf("Point 2 - Row: %d, Col: %d\n", point2.row, point2.col)
	fmt.Printf("Point 3 - Row: %d, Col: %d\n", point3.row, point3.col)
}
