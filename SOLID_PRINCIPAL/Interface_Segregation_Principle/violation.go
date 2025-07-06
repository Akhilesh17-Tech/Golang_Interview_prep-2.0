// The Interface Segregation Principle (ISP) is a key part of the SOLID design principles and states that "Clients should not be forced to depend on interfaces they do not use." This means that it's better to create smaller, more specific interfaces rather than large, monolithic ones that encompass a wide range of functionalities.

// Clients should not be forced to depend on interfaces they don't use."
package main

import "fmt"

// ISP Violation: A single large interface
type Printer interface {
	Print(document string)
	Scan(document string)
	Fax(document string)
}

// OldPrinter implements the Printer interface, but doesn't support all functionalities
type OldPrinter struct{}

func (o *OldPrinter) Print(document string) {
	fmt.Println("Printing document:", document)
}

func (o *OldPrinter) Scan(document string) {
	// OldPrinter doesn't support scanning, but still needs to implement this method
	fmt.Println("Scanning is not supported.")
}

func (o *OldPrinter) Fax(document string) {
	// OldPrinter doesn't support faxing, but still needs to implement this method
	fmt.Println("Faxing is not supported.")
}

func main() {
	printer := &OldPrinter{}
	printer.Print("Report")
	printer.Scan("Report") // Not supported
	printer.Fax("Report")  // Not supported
}
