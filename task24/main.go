package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func newPoint(x, y float64) *point {
	return &point{x, y}
}

func distance(a, b *point) float64 {
	return math.Sqrt(math.Pow((a.x-b.x), 2) + math.Pow((a.y-b.y), 2)) // Евклидово расстояние
}

func main() {
	point1 := newPoint(7, 10)
	point2 := newPoint(2, 5)
	fmt.Println(distance(point1, point2))
}
