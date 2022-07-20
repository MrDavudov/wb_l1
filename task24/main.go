package main

// Задания 24
// Разработать программу нахождения расстояния между двумя точками, 
// которые представлены в виде структуры Point с инкапсулированными 
// параметрами x,y и конструктором.

import (
	"fmt"
	"math"
)

type Point struct {
	x	float64
	y	float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x, 
		y: y,
	}
}

func (p Point) distanceTo(t Point) float64 {
	distX := p.x - t.y
	distY := p.y - t.x
	return math.Sqrt(distX*distX + distY*distY)
}

func main() {
	p1 := Point{12, 23}
	p2 := Point{10, 342}

	fmt.Printf("Дистанция от person1 до person2: %.2f", p1.distanceTo(p2))
}