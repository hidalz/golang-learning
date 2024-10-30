package main

func Sum(numbers []int) int {
	sum := 0

	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	numberLength := len(numbersToSum)

	sums := make([]int, numberLength)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
