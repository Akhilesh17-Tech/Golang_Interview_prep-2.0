package main

import "fmt"

func main() {
	// function with multi type return
	fmt.Println(multiFunc())

	// function with multi type arguments
	multiArgFunc(25, "kushwah", true)

	inc := displayNumber()

	// isolated number with closer new clouser
	inc1 := displayNumber()

	fmt.Println("number incremented : ", inc())
	fmt.Println("number incremented : ", inc())
	fmt.Println("number incremented : ", inc1())

	// varidic function can take multiple arguments at a time there is no limit same as rest operator in JavaScript
	fmt.Println("variadic 1", variadic(1, 2, 3, 4, 5, 6))
	fmt.Println("variadic 2", variadic(1, 2, 3))
	fmt.Println("variadic 3", variadic(1, 2, 3, 4, 5, 6, 7, 8, 9))

}

func multiFunc() (int, string) {
	return 20, "akhilesh"
}

func multiArgFunc(age int, surname string, marridge bool) {
	fmt.Println("age:", age, ",", "surname :", surname, ",", "marridge :", marridge)
}

// function assign to name
func functionAssignedToName() {
	fmt.Println("this is closer function")
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(2, 5))
}

// clouser function in go

func displayNumber() func() int {
	number := 0
	inc := func() int {
		number++
		return number
	}
	return inc
}

// variadic fuction

func variadic(args ...int) int {
	total := 0
	for _, v := range args { // Iterates over the arguments whatever the number.
		total += v
	}
	return total
}
