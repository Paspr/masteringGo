package main

/*
This package contains an implementation of a linked list for integers
*/

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (n LinkedList) Find(data int) *Node {
	// Finds a node by value
	currentNode := n.head
	for currentNode != nil {
		if currentNode.data == data {
			return currentNode
		}
		currentNode = currentNode.next
	}
	return nil
}

func (n *LinkedList) Clear() {
	// Clears linked list, sets head and tail to nil
	n.head = nil
	n.tail = nil
}

func (n LinkedList) Count() int {
	// Counts the nodes in a linked list
	if n.head == nil {
		return 0
	} else {
		currentNode := n.head
		length := 0
		for currentNode != nil {
			length++
			currentNode = currentNode.next
		}
		return length
	}
}

func (n *LinkedList) addInTail(item *Node) {
	// Adds a node at the end of linked list
	if n.head == nil {
		n.head = item
	} else {
		n.tail.next = item
	}
	n.tail = item
}

func (n LinkedList) printLinkedList() {
	// Prints out linked list
	currentNode := n.head
	for currentNode != nil {
		fmt.Print(" node data: ", currentNode.data, " node ptr: ", currentNode.next)
		currentNode = currentNode.next
	}
	fmt.Println()
}

// to do
// FindAll (slice append)
// Remove (bool)
// RemoveAll
// InsertAfter

func main() {
	var testList LinkedList
	testNode := Node{data: 10}
	testNode2 := Node{data: 20}
	testNode3 := Node{data: 30}

	testList.addInTail(&testNode)
	testList.addInTail(&testNode2)
	testList.addInTail(&testNode3)

	testList.printLinkedList()
	fmt.Println(testList.Count())

	if testList.head.data == 10 && testList.head.next.data == 20 && testList.head.next.next.data == 30 {
		fmt.Println("Linked list is OK")
	} else {
		fmt.Println("Linked list is incorrect!")
	}
}
