package main

import "fmt"

// Account is the base struct
type Account struct {
	accountNumber string
	accountHolder string
	balance       float64
}

// Deposit method for Account
func (a *Account) Deposit(amount float64) {
	a.balance += amount
	fmt.Printf("Deposited $%.2f into account %s\n", amount, a.accountNumber)
}

// Withdraw method for Account
func (a *Account) Withdraw(amount float64) error {
	if a.balance < amount {
		return fmt.Errorf("insufficient balance")
	}
	a.balance -= amount
	fmt.Printf("Withdrew $%.2f from account %s\n", amount, a.accountNumber)
	return nil
}

// CheckingAccount is a struct that embeds Account
type CheckingAccount struct {
	Account
	overdraftLimit float64
}

// SavingsAccount is a struct that embeds Account
type SavingsAccount struct {
	Account
	interestRate float64
}

// CreditCardAccount is a struct that embeds Account
type CreditCardAccount struct {
	Account
	creditLimit float64
}

func main() {
	// Create a CheckingAccount
	checkingAccount := CheckingAccount{
		Account: Account{
			accountNumber: "1234567890",
			accountHolder: "John Doe",
			balance:       1000.0,
		},
		overdraftLimit: 500.0,
	}

	// Deposit and withdraw money from the checking account
	checkingAccount.Deposit(500.0)
	checkingAccount.Withdraw(200.0)

	// Create a SavingsAccount
	savingsAccount := SavingsAccount{
		Account: Account{
			accountNumber: "9876543210",
			accountHolder: "Jane Doe",
			balance:       500.0,
		},
		interestRate: 0.05,
	}

	// Deposit and withdraw money from the savings account
	savingsAccount.Deposit(200.0)
	savingsAccount.Withdraw(100.0)

	// Create a CreditCardAccount
	creditCardAccount := CreditCardAccount{
		Account: Account{
			accountNumber: "1111111111",
			accountHolder: "Bob Smith",
			balance:       0.0,
		},
		creditLimit: 1000.0,
	}

	// Deposit and withdraw money from the credit card account
	creditCardAccount.Deposit(500.0)
	creditCardAccount.Withdraw(200.0)
}
