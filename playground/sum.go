package main

import "fmt"

var sum = func (v ...int)  int {
	var total int 

	for _, val := range v {
		total += val
	}

	return total
}

func main() {
	x := sum()
	y := sum(1,2)
	z := sum(1,2,3,4,5)

	fmt.Println(x, y, z)
}
