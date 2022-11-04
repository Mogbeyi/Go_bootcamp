package main

import "fmt"

func getItemCost(itemId int) (float64, error) {
	if itemId <= 0 {
		return 0.0, error("Invalid parameter, ItemId must be greater than 0")
	}

	return 12.94, error("")
}

func main() {
	fmt.Println(getItemCost(4))
}
