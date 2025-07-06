//High-level modules should not depend on low-level modules. Both high-level and low-level modules should depend on abstractions.

// In simpler terms, DIP suggests that:

// High-level modules (e.g., business logic) should not depend on low-level modules (e.g., database or file system).
// Instead, both high-level and low-level modules should depend on abstractions (e.g., interfaces).
package main

import "fmt"

// Abstraction (interface)
type OrderRepository interface {
	SaveOrder(order Order)
}

// High-level module (business logic)
type OrderService struct {
	repo OrderRepository
}

func (s *OrderService) PlaceOrder(order Order) {
	// Depend on abstraction (interface)
	s.repo.SaveOrder(order)
}

// Low-level module (database)
type Database struct{}

func (d *Database) SaveOrder(order Order) {
	// Database-specific implementation
	fmt.Println("Order saved:", order)
}

// Order type
type Order struct {
	ID   int
	Item string
}

func main() {
	// Create a Database instance and pass it to OrderService
	db := &Database{}
	orderService := &OrderService{repo: db}
	
	order := Order{ID: 1, Item: "Laptop"}
	orderService.PlaceOrder(order)
}
