package main

import (
	"fmt"
	"sync"
)

// A pointer in Go is a variable that holds the memory address of another variable. Pointers are useful for passing references to values and arrays within functions and can help in achieving efficient memory usage.

func main() {
	var i int = 10
	var ptr *int

	ptr = &i

	fmt.Println(ptr)  // this will print the address
	fmt.Println(*ptr) // dereferencing the value by using *, Dereferencing a pointer means accessing the value that the pointer points to.

	// we can modify the value of a variable through a pointer by dereferencing it and assigning a new value.
	*ptr = 32
	fmt.Println(i)

	// The zero value of a pointer in Go is nil. A nil pointer does not point to any memory location.
	var ptr2 *int
	fmt.Println(ptr2)

	// passing pointer to a function
	fmt.Println(pointerFunc(ptr))
	fmt.Println(i) // here automatically change
	pointerWithStruct()

	ponterToPoniter()
}

func pointerFunc(num *int) int {
	// modifying value here
	*num = 100
	return *num
}

func pointerWithStruct() {
	type Person struct {
		name string
		age  int
	}

	person := Person{name: "akhilesh kushwah", age: 25} // instance initialized

	var ptr *Person = &person

	fmt.Println(ptr.name)
	fmt.Println(ptr.age)

	// modifying the value
	ptr.name = "akash dayma"
	ptr.age = 29

	fmt.Println(ptr.name)
	fmt.Println(ptr.age)

	// concurrent programming
	var wg sync.WaitGroup
	data := 0
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go goRoutinesAndChannelsWithPointer(i, &wg, &data)
	}

	wg.Wait()
	fmt.Println("Final Data:", data)

	// largeDataStructure
	largeDataStructure()

	// pointerWithString
	str := "Hi Akhilesh"
	pointerWithString(&str)
	fmt.Println("Modified string : ", str)

	//pointer pracrtice call

	pointerBasicPractice()

}

// pointer to pointer when we need to pass reference to pointer itself
// Yes, you can have a pointer to a pointer in Go. This is useful when you need to reference a pointer itself.

func ponterToPoniter() {
	var a int = 50
	var ptr *int = &a
	var ptrToPtr **int = &ptr
	fmt.Println(ptr)        // Prints the memory address of 'a'
	fmt.Println(*ptr)       // Prints 42 dereferencing that *ptr points to a
	fmt.Println(ptrToPtr)   // Prints the memory address of 'ptr'
	fmt.Println(**ptrToPtr) // Prints 42 dereferencing that **ptrToPtr to ptr

	fmt.Println(*ptr == a) //true
}

// Pointers when working with go routines and channels

func goRoutinesAndChannelsWithPointer(i int, wg *sync.WaitGroup, data *int) {
	defer wg.Done()
	*data = *data + i
	fmt.Printf("Worker %d, Data: %d\n", i, *data)
}

// Avoiding Copying Large Data Structures

type LargeStruct struct {
	Data [100]int
}

func largeDataStructure() {
	var ls LargeStruct
	modifyStruct(&ls)
	fmt.Println("largeDataStructure :", ls.Data[0]) // Prints 100

}

func modifyStruct(ls *LargeStruct) {
	ls.Data[0] = 100
}

func pointerWithString(str *string) {
	*str = "Hello Akhilesh Ji Kushwah"
}

func pointerBasicPractice() {

	fmt.Println("Pointer basic calling .............")
	var a int = 100
	var ptr *int = &a
	var ptr1 *int = ptr

	*ptr = 1000
	fmt.Println(*ptr, *ptr1)
	fmt.Println("Pointer basic calling end .............")

}
