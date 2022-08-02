package main

import "fmt"

type game struct {
	name string
	price float64
}

func (g game) print() {
	fmt.Println(g.name, ": ", g.price)
}