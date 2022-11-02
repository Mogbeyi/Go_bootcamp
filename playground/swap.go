package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	x := 4
	y := 5
	x, y = swap(x, y)

	fmt.Println(x, y)
}

