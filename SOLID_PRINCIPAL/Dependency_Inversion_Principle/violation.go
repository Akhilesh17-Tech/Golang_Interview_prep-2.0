package main

import "fmt"

// High-level module (business logic)
type OrderService struct{}

func (s *OrderService) PlaceOrder(order Order) {
	// Depend on low-level module (database)
	db := &Database{}
	db.SaveOrder(order)
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
	orderService := &OrderService{}
	order := Order{ID: 1, Item: "Laptop"}
	orderService.PlaceOrder(order)
}
