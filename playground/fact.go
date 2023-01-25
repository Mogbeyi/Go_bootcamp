package main

import "fmt"

func factorial(n int) int {
	var res = 1

	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}

func main() {
	fact := factorial(5)

	fmt.Println(fact)
}

