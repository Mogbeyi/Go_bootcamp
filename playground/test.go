package main

import "fmt"

func num(v ...int) {
	fmt.Println(v)
}

func main() {
	num(1,2,3,4,5)
}
