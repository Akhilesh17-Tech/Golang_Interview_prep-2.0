// OCP compliant

// Now, adding a new payment method just requires creating a new struct that implements the PaymentMethod interface without touching the existing ProcessPayment function.

// PaymentMethod interface defines a method for processing payments
package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64)
}

// CreditPayment struct represents a credit card payment method
type CreditPayment struct{}

// Pay processes the credit card payment
func (c *CreditPayment) Pay(amount float64) {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
}

// PayPalPayment struct represents a PayPal payment method
type PayPalPayment struct{}

// Pay processes the PayPal payment
func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
}

func ProcessPayment(method PaymentMethod, amount float64) {
	method.Pay(amount)
}

func main() {
	// Creating different payment methods
	creditPayment := &CreditPayment{}
	payPalPayment := &PayPalPayment{}

	// Processing payments with different methods
	ProcessPayment(creditPayment, 100.00) // Outputs: Processing credit card payment of $100.00
	ProcessPayment(payPalPayment, 150.50) // Outputs: Processing PayPal payment of $150.50

}
