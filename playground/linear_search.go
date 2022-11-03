package main

import "fmt"

func linearSearch(data []int, target int) bool {
	for _, x := range data {
		if x == target {
			return true
		}
	}

	return false
}

func main() {
	items := []int{95,78,46,58,45,86,99,251,320}
    fmt.Println(linearSearch(items,58))
}
