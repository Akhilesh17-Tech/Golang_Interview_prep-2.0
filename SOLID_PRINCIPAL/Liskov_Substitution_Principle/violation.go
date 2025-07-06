// LSP says that if you have a function that works with a base class (or interface), it should also work with any subclass (or implementation of that interface) without causing errors or unexpected behavior.

// In simpler words:

// If you have a program that works with a parent class, it should still work if you replace the parent class with any of its child classes.

package main

import "fmt"

// Bird struct
type Bird struct{}

func (b *Bird) Fly() {
	fmt.Println("Flying!")
}

// Penguin struct, a type of Bird
type Penguin struct {
	Bird
}

func main() {
	// Create a bird
	bird := &Bird{}
	bird.Fly() // Outputs: Flying!

	// Create a penguin
	penguin := &Penguin{}
	penguin.Fly() // Uh-oh, penguins can't fly, but this will still output: Flying!
}
