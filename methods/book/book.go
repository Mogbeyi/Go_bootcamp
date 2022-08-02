package main

import "fmt"

type book struct {
	title string
	price float64
}

func (b book) print() {
	fmt.Println(b.title, ": ", b.price)
}