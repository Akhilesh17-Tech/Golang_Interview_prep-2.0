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

func PrintList(node *Node) {
	for node != nil {
		fmt.Printf("%d -> ", node.Value)
		node = node.Next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)

	PrintList(list.Head)
	newNode := RemoveNthNodeFromLast(list.Head, 6)
	PrintList(newNode)

}

//  1 -> 2 -> 3 -> 4 -> 5 -> 6 -> nil  // length = 6
// remove nth node from last if n = 2 then remove 5 and linked list should be 1 -> 2 -> 3 -> 4 -> 6 -> nil

func RemoveNthNodeFromLast(head *Node, n int) *Node {
	if head == nil {
		return head
	}

	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}
	target := length - n
	if target == 0 {
		return head.Next
	}

	current = head
	for i := 0; i < target-1; i++ {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	}

	return head
}
