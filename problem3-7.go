package main

import (
	"fmt"
)

type Point struct {
	row int
	col int
}

func IncPtr(p *Point) {
	p.row++
	p.col++
}

func IncNoPtr(p Point) {
	p.row++
	p.col++
}

func IncNewPtr(p **Point) {
	*p = &Point{row: (*p).row + 1, col: (*p).col + 1}
}

func main() {

	point1 := Point{row: 1, col: 1}
	point2 := point1
	point3 := &point1

	IncPtr(point3)

	IncNoPtr(point2)

	var newPoint *Point = &point1
	IncNewPtr(&newPoint)

	fmt.Printf("Point1: (%d, %d)\n", point1.row, point1.col)
	fmt.Printf("Point2: (%d, %d)\n", point2.row, point2.col)
	fmt.Printf("New Point: (%d, %d)\n", newPoint.row, newPoint.col)
}
