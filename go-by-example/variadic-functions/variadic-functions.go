package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// Convert []int to []any
	anyNums := make([]any, len(nums))
	for i, v := range nums {
		anyNums[i] = v
	}

	fmt.Println(anyNums...) // Here, I can't put "string", value since it would be too many values. Would need to include it

	anynumsdirect := []any{1, 2, 3, 4}
	fmt.Println("anynumsdirect: ", anynumsdirect)
}
