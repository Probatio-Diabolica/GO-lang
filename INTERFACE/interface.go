package main

import (
	"fmt"
	"math"
)

// Abser : defining an interface
type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

type MyFLoat float64

func (f MyFLoat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser
	v := Vertex{3, 4}
	f := MyFLoat(-7)

	a = v

	fmt.Println(a.Abs())

	a = f
	fmt.Println(a.Abs())
}
