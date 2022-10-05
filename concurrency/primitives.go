package main

import (
	"fmt"
	"time"
)

func someFunc(s string) {
	fmt.Println(s)
}

func main() {
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")

	time.Sleep(time.Second * 2)

	fmt.Println("Hello world")
}
