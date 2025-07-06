package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (a *BankAccount) Deposit(amount int) {
	a.mu.Lock()
	a.balance += amount
	a.mu.Unlock()
}

func (a *BankAccount) Withdraw(amount int) {
	a.mu.Lock()
	a.balance -= amount
	a.mu.Unlock()
}

func (a *BankAccount) GetBalance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func main() {
	account := &BankAccount{balance: 100}

	// var wg sync.WaitGroup
	// wg.Add(2) // We have 2 goroutines to wait for 



	go func() {
		account.Deposit(50)
		// wg.Done()
		// time.Sleep(10 * time.Second)
	}()

	go func() {
		account.Withdraw(20)
		// wg.Done()
		// time.Sleep(10 * time.Second)
	}()

	time.Sleep(10 * time.Second)
	// wg.Wait() // Wait for both goroutines to finish

	fmt.Println(account.GetBalance()) // Output: 130
}




