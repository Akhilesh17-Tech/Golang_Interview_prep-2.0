package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents the entire linked list
type LinkedList struct {
	Head *Node
	Size int
}

func (list *LinkedList) Append(value int) {
	newNode := &Node{Value: value}
	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	list.Size++
}

func (list *LinkedList) Prepend(value int) {
	newNode := &Node{Value: value}
	if list.Head == nil {
		list.Head = newNode
	} else {
		newNode.Next = list.Head
		list.Head = newNode
	}
	list.Size++
}

func (list *LinkedList) Delete(value int) {
	if list.Head == nil {
		fmt.Println("Empty list can't delete")
	}
	if list.Head.Value == value {
		list.Head = list.Head.Next
		return
	}
	current := list.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
	list.Size--
}

func (list *LinkedList) Search(value int) *Node {
	if list.Head == nil {
		return nil
	}
	current := list.Head
	for current.Next != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (list *LinkedList) Print() {
	if list.Head == nil {
		fmt.Println("List is empty nothing to Print")
	} else {
		current := list.Head
		for current != nil {
			fmt.Printf("%d -> ", current.Value)
			current = current.Next
		}
		fmt.Println("nil")
	}

}

func main() {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Prepend(10)
	list.Prepend(11)
	list.Prepend(12)
	list.Print()
	node := list.Search(1)
	if node == nil {
		fmt.Println("Element not found in the Linked List")
	} else {
		fmt.Println("Element found in the Linked List", node.Value)
	}
	fmt.Println("size of the linked list = ", list.Size)

	node1 := ReverseLinkedList(list.Head)
	PrintList(node1)
	node2 := ReverseLinkedListUsingStack(node1)
	PrintList(node2)

}

// recursive approach
func ReverseLinkedList(node *Node) *Node {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := ReverseLinkedList(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

func PrintList(node *Node) {
	for node != nil {
		fmt.Printf("%d -> ", node.Value)
		node = node.Next
	}
	fmt.Println("nil")
}

// using stack

func ReverseLinkedListUsingStack(node *Node) *Node {
	if node == nil || node.Next == nil {
		return node
	}

	// Stack to hold the nodes
	stack := []*Node{}

	// Push all nodes onto the stack
	current := node
	for current != nil {
		stack = append(stack, current)
		current = current.Next
	}

	// The last node becomes the new head
	newHead := stack[len(stack)-1]

	// Re-link the nodes in reverse order
	for i := len(stack) - 1; i > 0; i-- {
		stack[i].Next = stack[i-1]
	}

	// The original head becomes the new tail, so its next should be nil
	stack[0].Next = nil

	return newHead

}
