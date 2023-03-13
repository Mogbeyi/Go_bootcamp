package main

import (
	"fmt"
	"math"
)
 
type Circle struct {
	x, y, r float64
}


func main() {
	c := Circle{x: 4, y: 9, r: 2}

	fmt.Println(calcArea(c))
}

func calcArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}
