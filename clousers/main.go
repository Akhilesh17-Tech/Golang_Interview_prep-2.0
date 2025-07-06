//A closure in Go is a function value that references variables from outside its body. The function may access and modify the variables referenced from outside its body, even after the surrounding function has returned.

// Closures are useful in Go for several reasons, including encapsulating behavior, maintaining state between function calls, and creating more concise and readable code.

// Definition
// A closure is created when a function is defined within another function and references variables from the outer function. The inner function "closes over" the variables from the outer function.

package main

import "fmt"

func Counter() (func() int, func() int) {
	counter := 0 // encapsulated counter variable cannot be accessed directly
	increment := func() int {
		counter++
		return counter
	}
	decrement := func() int {
		counter--
		return counter
	}

	return increment, decrement
}

func main() {
	increment, decrement := Counter()
	fmt.Println(increment()) // prints 1
	fmt.Println(increment()) // prints 2
	fmt.Println(decrement()) // prints 1
}
