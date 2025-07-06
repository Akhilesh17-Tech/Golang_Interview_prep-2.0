package main

import "fmt"

// Flyable interface
type Flyable interface {
	Fly()
}

// Bird struct
type Bird struct{}

func (b *Bird) Chirp() {
	fmt.Println("Chirping!")
}

// Sparrow struct, a type of bird that can fly
type Sparrow struct {
	Bird
}

func (s *Sparrow) Fly() {
	fmt.Println("Flying!")
}

// Penguin struct, a type of bird that cannot fly
type Penguin struct {
	Bird
}

func (p *Penguin) Swim() {
	fmt.Println("Swimming!")
}

func main() {
	// Create a sparrow
	sparrow := &Sparrow{}
	sparrow.Fly()   // Outputs: Flying!
	sparrow.Chirp() // Outputs: Chirping!

	// Create a penguin
	penguin := &Penguin{}
	penguin.Swim()  // Outputs: Swimming!
	penguin.Chirp() // Outputs: Chirping!
}
