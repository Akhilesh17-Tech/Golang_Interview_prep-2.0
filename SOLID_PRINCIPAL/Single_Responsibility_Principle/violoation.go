// Single Responsibility Principle (SRP)
// What It Means: A piece of code (like a function, method, or class) should only do one thing. If you find yourself changing the same piece of code for different reasons, it's probably violating SRP.

// Example Explanation:
// Imagine you have a User struct that represents a user in your application. Now, let's say you also want to handle saving the user to a database and sending a welcome email when a new user is created. If you put both saving and emailing logic in the same function, this function is now responsible for two different things: database interaction and email sending.
// If the way you save users changes, you'll modify this function. If the way you send emails changes, you'll modify this function again. This breaks SRP because one function handles multiple responsibilities.

package main

import (
	"fmt"
)

// SRP Violation User struct is tightly coupled
type User struct {
	Username string
	Email    string
}

// Handles both user management and email notifications
func (u *User) SaveUser() {
	// Save user to database
	fmt.Println("Saving user:", u.Username)
}

func (u *User) SendWelcomeEmail() {
	// Send a welcome email
	fmt.Println("Sending welcome email to:", u.Email)
}

func main() {
	// Create a new user
	user := &User{Username: "john_doe", Email: "john.doe@example.com"}

	// Save the user
	user.SaveUser()

	// Send a welcome email
	user.SendWelcomeEmail()
}
