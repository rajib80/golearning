package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

func (anySquare square) area() float64 {
	return anySquare.side * anySquare.side
}

type circle struct {
	radius float64
}

func (anycircle circle) area() float64 {
	return math.Pi * (anycircle.radius * anycircle.radius)
}

func info(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	aSquare := square{5}
	info(aSquare)

	aCircle := circle{6}
	info(aCircle)
}
