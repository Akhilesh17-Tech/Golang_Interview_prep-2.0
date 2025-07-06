package main

import "fmt"

// interface and its work
// defer
// go rutine
// channel
// pointer
//

// interface define : interface is collection of method singature and the interface is abstract that means we are not allowed to create instance of interface, but we are allowed to create variable of interface

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct type first letter capital means it is exported and available outside the package
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the interface methods for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Define another struct type
type Circle struct {
	Radius float64
}

// Implement the interface methods for Circle πr^2
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
	// area of circle ?
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}



type Car struct {
	Name string
	Horsepower int
	Type string 
}

func main() {
	// Create instances of Rectangle and Circle
	rect := Rectangle{Width: 5, Height: 4}
	circ := Circle{Radius: 3}

	// Create a slice of Shape interface
	shapes := []Shape{rect, circ}

	// Iterate through the shapes and print their area and perimeter
	for _, shape := range shapes {
		fmt.Printf("Shape: %T, Area: %.2f, Perimeter: %.2f\n", shape, shape.Area(), shape.Perimeter())
	}

	// struct function call
	strucExplain()
}

func strucExplain() {
	// Define a struct type

	type Person struct {
		name       string
		age        int
		ismarridge bool
		salary     float32
	}

	//JSON tags are annotations used in struct fields to specify the key name in JSON encoding and decoding. They allow customization of field names when a struct is serialized to or deserialized from JSON format. For example, the Mobile field will be encoded as mobile in the resulting JSON.
	type GenericOtpProvider struct {
		Mobile   string `json:"mobile"`
		Count    int    `json:"count"`
		Provider string `json:"provider"`
		UseCase  string `json:"usecase"`
	}

	Person1 := Person{name: "akhilesh", age: 25, ismarridge: false, salary: 100000.310}

	fmt.Println(Person1) // will print the entire struct

	fmt.Println(Person1.age) // will print the specific value based on key

}

// Applications of struct
// Data Organization: Structs are ideal for grouping related data together, making it easier to manage and understand.
// Modeling Real-world Entities: Structs can represent complex entities like a Person, Car, Book, etc., in your program.
// Encapsulation: Structs can be used to encapsulate data and related methods, similar to objects in object-oriented programming.
// Serialization: Structs are commonly used for encoding data into JSON, XML, or other formats for storage or transmission.
