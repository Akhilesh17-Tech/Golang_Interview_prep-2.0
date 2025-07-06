package main

import (
	"fmt"
)

func main() {
	// variable declaration using var keyword with explicit type
	var name string = "akhilesh"
	var age int = 25
	var isMarried bool = true
	var salary float32 = 100000.00

	// variable declaration using const key explicit type
	const PI float32 = 3.14

	// := type not need to define short way
	fullName := "akhilesh kushwah"

	// type inference automatically type will sense based on value
	var address = "indore"

	// multiple variables using var
	var a, b, c int

	a, b, c = 10, 20, 30

	fmt.Println(name, fullName, age, isMarried, salary, address, a, b, c)

	// default value
	var o int
	var p string
	var q bool
	var r float32
	fmt.Println("o:", o) // o: 0
	fmt.Println("p:", p) // p: ""
	fmt.Println("q:", q) // q: false
	fmt.Println("r:", r) // r: 0

	// var name string

	var surname string
	surname = "kushwah"

	fmt.Println(surname)

}
