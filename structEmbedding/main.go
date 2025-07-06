package main

import "fmt"

func main() {

	// Embedding in Go allows structs and interfaces to compose behavior and fields from other structs and interfaces. This approach is preferred over traditional inheritance, providing a more flexible and maintainable way to structure code.

	// struct embedding

	type Person struct {
		name string
		age  int
	}

	type Employee struct {
		Person
		employeId string
		salary    float32
	}

	emp := Employee{
		Person: Person{
			name: "Akhilesh kushwah",
			age:  25,
		},
		employeId: "123",
		salary:    100000,
	}

	// Accessing fields directly
	fmt.Println("Name:", emp.name)
	fmt.Println("Age:", emp.age)
	fmt.Println("Employee ID:", emp.employeId)
	fmt.Println("Position:", emp.salary)

	p := Persons{name: "akhilesh ji"}

	// Use the methods from both interfaces
	var g Greeter = p

	g.greet()
	g.speak()

}

// interface embedding

type Speaker interface {
	speak()
}

type Greeter interface {
	Speaker // embedded
	greet()
}

type Persons struct {
	name string
}

func (p Persons) greet() {
	fmt.Println("Hello", p.name)
}
func (p Persons) speak() {
	fmt.Println("I am speaking mannn !!!!!!!!")
	// var arr = [5]int{}

	// mp1 := make(map[string]int)
	// mp := map[string]int{}
}
