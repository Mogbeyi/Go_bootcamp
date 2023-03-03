package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result)
}

func productExceptSelf(nums []int) []int {
	prefixResult := calcPrefix(nums)
	suffixResult := calcSuffix(prefixResult, nums)
	return suffixResult
}

func calcPrefix(nums []int) []int {
	result := make([]int, len(nums))
	productSum := 1

	for i := range nums {
		result[i] = productSum
		productSum *= nums[i]
	}

	return result
}

func calcSuffix(prefixResult, nums []int) []int {
	suffixResult := make([]int, len(nums))
	productSum := 1

	for i := len(nums) - 1; i >= 0; i-- {
		suffixResult[i] = productSum * prefixResult[i]
		productSum *= nums[i]
	}

	return suffixResult
}
