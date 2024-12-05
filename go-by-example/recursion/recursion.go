package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(60))

	// Anonymous functions can also be recursive, but this requires explicitly declaring a
	// variable with var to store the function before it’s defined.
	var fib func(n int) int

	/*
		The reason you need to declare the header of the function using var fib func(n int) int
		before defining the function is to allow the function to refer to itself recursively. In Go,
		functions are not fully defined until their entire body is parsed. By declaring the function
		type first, you create a variable that can be assigned the function later, allowing the
		function to call itself within its body.

		Declaration:
		var fib func(n int) int declares a variable fib of type func(n int) int.
		This tells the compiler that fib will hold a function that takes an int as an argument
		and returns an int.

		Definition:
		fib = func(n int) int { ... } defines the function and assigns it to the variable fib.
		Now, within the function body, fib can be used to refer to itself, enabling recursion.

		Why This is Necessary:
		Without the initial declaration, the function would not be able to refer to itself because
		it would not be fully defined at the time of its first use within its own body.
		This would result in a compilation error.


	*/
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
