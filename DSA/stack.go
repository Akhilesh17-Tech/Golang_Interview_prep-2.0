package main

import "fmt"

// Stack represents a stack data structure
type Stack struct {
	elements []int
}

// Push adds an element to the top of the stack
func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

// Pop removes and returns the top element of the stack
// It returns an error if the stack is empty
func (s *Stack) Pop() (int, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	// Get the top element
	top := s.elements[len(s.elements)-1]
	// Remove the top element
	s.elements = s.elements[:len(s.elements)-1]
	return top, nil
}

// Peek returns the top element without removing it
// It returns an error if the stack is empty
func (s *Stack) Peek() (int, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.elements)
}

func main() {
	stack := Stack{elements: []int{1, 2, 4, 5, 5, 7, 8, 9}}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack size:", stack.Size()) // Output: Stack size: 3

	top, err := stack.Peek()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Top element:", top) // Output: Top element: 3
	}

	popElement, err := stack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Popped element:", popElement) // Output: Popped element: 3
	}

	fmt.Println("Stack size after pop:", stack.Size()) // Output: Stack size after pop: 2
}
