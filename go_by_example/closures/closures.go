package main

import "fmt"

// intSeq returns an anonymous function that is a closure
func intSeq() func() int {
	i := 0              // Variable in the outer function's scope
	return func() int { // Anonymous function
		i++ // The anonymous function captures and modifies 'i'
		return i
	}
}

// intSeq returns a named function that is a closure. It basically works with anonymous and not anonymous functions.
func intSeqNamed() func() int {
	i := 0 // Variable in the outer function's scope

	// Named function
	increment := func() int {
		i++ // The named function captures and modifies 'i'
		return i
	}

	return increment // Return the named function
}

func main() {
	nextInt := intSeq() // nextInt is a closure that captures 'i'

	fmt.Println(nextInt()) // Prints 1
	fmt.Println(nextInt()) // Prints 2
	fmt.Println(nextInt()) // Prints 3

	newInts := intSeq()    // newInts is a new closure with a new 'i'
	fmt.Println(newInts()) // Prints 1
}
