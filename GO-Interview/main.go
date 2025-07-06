package main

import "fmt"

func main() {
	// var x,y,z = 10, 1.5, "Akhilesh"
	map1 := make(map[string]bool, 10)

	map1["interview"] = true
	map1["bit"] = false

	map2 := make(map[string]bool, 10) // why make -> iniialize initial capacity

	for i, v := range map1 {
		map2[i] = v
	}

	map2["bit"] = true

	fmt.Println(map1, map2)

	// inheritance call
	Inheritance()

	//polymorphism call
	polymorphism()

}

type Person struct {
	name      string
	age       int
	isMarride bool
}

type Company struct {
	Person      //embedding also known
	companyName string
	address     string
	contact     int
}

func (p Person) greet() {
	fmt.Println("Hello : ", p.name, p.age)
}

func (c Company) showDetails() {
	fmt.Println("Company name :", c.companyName, "company address :", c.address)
}

func Inheritance() {

	var p = Person{name: "akhilesh", age: 25}
	p.isMarride = true
	fmt.Println(p.isMarride, "ismaridged")
	var c = Company{Person: p, companyName: "RWS", address: "Indore", contact: 9981806164}
	c.greet()
	c.showDetails()
	// fmt.Println(c)

}

// polymorphism

type animal interface {
	speak() string
}

type cat struct{}

func (c cat) speak() string {
	return "meow"
}

type dog struct{}

func (d dog) speak() string {
	return "bark"
}

func makeSound(a animal) {
	fmt.Println(a.speak())
}

func polymorphism() {
	var a animal
	a = cat{}
	makeSound(a)
	// fmt.Println(a.speak())
	a = dog{}
	makeSound(a)
	// fmt.Println(a.speak())
}
