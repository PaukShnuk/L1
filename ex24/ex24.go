package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

// New - функция-конструктор
func New(x int, y int) Point {
	return Point{x, y}
}

// squareCalculate - функция вычисления расстояния
func squareCalculate(p1 Point, p2 Point) float64 {
	k1 := math.Abs(float64(p1.x - p2.x))
	k2 := math.Abs(float64(p1.y - p2.y))
	return math.Sqrt(math.Pow(k1, 2) + math.Pow(k2, 2))
}

func main() {

	p1 := New(-5, 1)
	p2 := New(3, -1)
	fmt.Println(squareCalculate(p2, p1))
}
