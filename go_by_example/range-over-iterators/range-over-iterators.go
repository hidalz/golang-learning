// TODO: Review
package main

import (
	"fmt"
	"iter"
	"slices"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) All() iter.Seq[T] { // Function signature, it is just a closure that returns an iterator. Therefore the (an) iterator is the closure that serves to maintain state variables for the inside function.
	return func(yield func(T) bool) { // ANonymous function (closure) that takes yet a third function that returns a bool.
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
} // TODO: review iterators and the closure-based iterator pattern in Go. Still don't get it fully..
/*
The code you provided demonstrates the closure-based iterator pattern in Go. This pattern allows you to define an iterator function that can be used in a for loop to iterate over a collection of values.

In the code snippet, the All() method is defined on the List type. It returns a closure that takes a yield function as a parameter. This closure represents the iterator. The yield function is a callback that will be called for each element in the list.

Inside the closure, a for loop is used to iterate over the elements of the list. For each element, the yield function is called with the element's value. If the yield function returns false, the loop is terminated early using the return statement.

Now, let's address your questions:

Why do you need a closure for this pattern? The closure is used to capture the state of the iterator function. It allows the iterator to maintain its position in the iteration and remember the context in which it was created. In this case, the closure captures the lst variable, which represents the list being iterated over. This ensures that the iterator operates on the correct list even if the List object is modified after the iterator is created.

How would the yield ever return true? I don't see a return true. In this pattern, the yield function is responsible for controlling the iteration. It is up to the caller of the iterator to decide when to stop the iteration by returning false from the yield function. The return true statement is not explicitly written in the code because it is the default behavior when the yield function is not called with false. So, if the yield function is not called or called with true, the iteration continues.

How is the yield being passed to the iterator in the for loops? In the for loop, range lst.All() is used to obtain the iterator from the All() method. The iterator is a closure that takes the yield function as a parameter. The range keyword automatically calls the iterator and assigns the value returned by the yield function to the loop variable (e in this case). The loop continues until the iterator returns false from the yield function.

I hope this clarifies the closure-based iterator pattern and how it is used in the provided code. Let me know if you have any further questions or if there's anything else I can assist you with!
*/
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}

	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break
		}
		/*
			In this code, there is an if statement that checks if the value of n is greater than
			or equal to 10. If this condition is true, the break statement is executed. This causes
			the loop to exit immediately, skipping any remaining iterations.

			This code snippet is inside a loop that iterates over a collection using the range
			keyword. The range keyword is used to iterate over elements of an array, slice,
			string, map, or channel. In each iteration, the iterator's yield function is called to
			provide the next value from the collection.

			When the loop encounters the break statement, it exits the loop and the iterator's yield
			function returns false. This indicates to the loop that there are no more values to
			iterate over, and the loop terminates.
		*/
		fmt.Println(n)
	}
}
