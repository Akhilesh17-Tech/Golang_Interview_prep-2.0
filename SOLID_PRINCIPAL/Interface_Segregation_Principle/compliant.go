package main

import "fmt"

// Smaller, specific interfaces
type Printer interface {
	Print(document string)
}

type Scanner interface {
	Scan(document string)
}

type Faxer interface {
	Fax(document string)
}

// OldPrinter only implements the Printer interface
type OldPrinter struct{}

func (o *OldPrinter) Print(document string) {
	fmt.Println("Printing document:", document)
}

// ModernPrinter implements all interfaces, since it supports all functionalities
type ModernPrinter struct{}

func (m *ModernPrinter) Print(document string) {
	fmt.Println("Printing document:", document)
}

func (m *ModernPrinter) Scan(document string) {
	fmt.Println("Scanning document:", document)
}

func (m *ModernPrinter) Fax(document string) {
	fmt.Println("Faxing document:", document)
}

func main() {
	oldPrinter := &OldPrinter{}
	oldPrinter.Print("Old Report")

	modernPrinter := &ModernPrinter{}
	modernPrinter.Print("Modern Report")
	modernPrinter.Scan("Modern Report")
	modernPrinter.Fax("Modern Report")
}
