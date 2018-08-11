package main

import (
	"fmt"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	height float64
	width  float64
}

type triangle struct {
	side1  float64
	side2  float64
	side3  float64
	height float64
}

func (anyRectangle rectangle) area() float64 {
	return anyRectangle.height * anyRectangle.width
}

func (anyRectangle rectangle) perimeter() float64 {
	return 2 * (anyRectangle.width + anyRectangle.height)
}

func (anyTriangle triangle) area() float64 {
	return 0.5 * anyTriangle.side3 * anyTriangle.height
}

func (anyTriangle triangle) perimeter() float64 {
	return anyTriangle.side1 + anyTriangle.side2 + anyTriangle.side3
}

func measure(anyGeometry geometry) {
	fmt.Println(anyGeometry)
	fmt.Println(anyGeometry.area())
	fmt.Println(anyGeometry.perimeter())
}

func main() {
	aRectangle := rectangle{width: 5, height: 9}
	aTriangle := triangle{side1: 8, side2: 8, side3: 8, height: 6.93}

	measure(aRectangle)
	measure(aTriangle)
}
