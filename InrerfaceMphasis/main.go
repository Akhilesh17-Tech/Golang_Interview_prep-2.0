package main

import "fmt"

// Employee interface with a method to calculate salary
type Employee interface {
	CalculateSalary() float64
}

// Permanent struct representing a permanent employee
type Permanent struct {
	empId    int
	basicPay float64
}

// Implement CalculateSalary method for Permanent
func (p Permanent) CalculateSalary() float64 {
	return p.basicPay
}

// Contract struct representing a contract employee
type Contract struct {
	empId    int
	basicPay float64
}

// Implement CalculateSalary method for Contract
func (c Contract) CalculateSalary() float64 {
	return c.basicPay
}

// Freelancer struct representing a freelancer
type Freelancer struct {
	empId      int
	hourlyRate float64
	hoursWorked float64
}

// Implement CalculateSalary method for Freelancer
func (f Freelancer) CalculateSalary() float64 {
	return f.hourlyRate * f.hoursWorked
}

// CalculateTotalExpenses function to calculate the total salary expense for the company
func CalculateTotalExpenses(employees []Employee) float64 {
	total := 0.0
	for _, employee := range employees {
		total += employee.CalculateSalary()
	}
	return total
}

func main() {
	// Creating instances of different employee types
	p1 := Permanent{empId: 1, basicPay: 50000}
	c1 := Contract{empId: 3, basicPay: 40000}
	f1 := Freelancer{empId: 4, hourlyRate: 50, hoursWorked: 100}

	// Slice of Employee interface to store all types of employees
	employees := []Employee{p1, c1, f1}

	// Calculate the total expenses
	totalExpenses := CalculateTotalExpenses(employees)
	fmt.Printf("Total Expenses of the company: %.2f\n", totalExpenses)
}
