package main

import "fmt"

// SRP compliant

type User struct {
	Username string
	Email    string
}

// UserRepository handles saving users to the database
type UserRepository struct{}

// SaveUser saves the user to the database
func (r *UserRepository) SaveUser(user *User) {
	// Simulate saving to a database
	fmt.Println("Saving user to database:", user.Username)
}

// EmailService handles sending emails
type EmailService struct{}

// SendWelcomeEmail sends a welcome email to the user
func (e *EmailService) SendWelcomeEmail(user *User) {
	// Simulate sending an email
	fmt.Println("Sending welcome email to:", user.Email)
}

// UserController orchestrates the operations using UserRepository and EmailService
type UserController struct {
	userRepo     *UserRepository
	emailService *EmailService
}

// RegisterUser handles the user registration process
func (uc *UserController) RegisterUser(user *User) {
	uc.userRepo.SaveUser(user)
	uc.emailService.SendWelcomeEmail(user)
}

func main() {
	user := &User{
		Username: "johndoe",
		Email:    "john@example.com",
	}

	userRepo := &UserRepository{}
	emailService := &EmailService{}

	userController := &UserController{
		userRepo:     userRepo,
		emailService: emailService,
	}

	// Register the user, which saves them to the database and sends a welcome email
	userController.RegisterUser(user)
}
