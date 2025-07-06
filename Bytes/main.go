// In Go, the byte type is an alias for uint8, which means it is an 8-bit unsigned integer. The byte type is typically used to represent a byte of data, which is a common data unit in programming, especially when dealing with raw binary data or strings.
// The byte type is not a distinct type in Go, but rather an alias for uint8.
// This means that any operation that can be performed on a uint8 can also be performed on a byte, and vice versa.
// String Manipulation: Since strings in Go are sequences of bytes, you often work with byte slices when manipulating strings.
// Binary Data: When dealing with raw binary data, such as reading from or writing to files, network communication, or encoding/decoding operations.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data := []byte("Akhilesh")

	fmt.Println(data)

	dataInt := []byte{1, 2, 3, 4, 5, 6, 7}
	for _, v := range dataInt {
		fmt.Println(v)
	}

	// using byte writing to a file and reading from file
	FileReadingAndWriting()
}

// Reading and Writing Files

func FileReadingAndWriting() {
	// Writing to a file
	data := []byte("Hello, Go!")
	err := ioutil.WriteFile("example.txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Reading from a file
	readData, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(readData)) // Output: Hello, Go!
}
