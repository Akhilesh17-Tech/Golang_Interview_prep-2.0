//What It Means: Your code should be open for adding new features or extending functionality but closed to modifying existing code.

// Example Explanation:
// Suppose you're writing a payment system. Initially, it supports credit card payments. Later, you want to add PayPal. If you have to change the existing code to add PayPal, it violates OCP.

// Every time you add a new payment method, you must modify the ProcessPayment function.
package main

import "fmt"

// OCP violation
func ProcessPayment1(method string) {
	if method == "credit" {
		// Process credit card payment
		fmt.Println("payment made via credit card.....")
	} else if method == "paypal" {
		// Process PayPal payment
		fmt.Println("payment made via paypal .....")

	}
}

func main() {
	// Suppose new payment methods are added, you need to modify the ProcessPayment1 function
	ProcessPayment1("credit") // Process credit card payment
	ProcessPayment1("paypal") // Process PayPal payment
}
