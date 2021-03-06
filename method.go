package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x float64
	y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	fmt.Println(variable())
}

func variable() int {
	a := 2
	b := 4

	return a * b
}
