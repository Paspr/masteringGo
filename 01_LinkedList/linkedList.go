package main

/*
This package contains a linked list data structure implementation for integers
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

func (n *LinkedList) addInTail(item *Node) {
	// adds a node at the of linked list
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

func main() {
	var testList LinkedList
	testNode := Node{data: 10}
	testNode2 := Node{data: 20}
	testNode3 := Node{data: 30}

	testList.addInTail(&testNode)
	testList.addInTail(&testNode2)
	testList.addInTail(&testNode3)

	testList.printLinkedList()

	if testList.head.data == 10 && testList.head.next.data == 20 && testList.head.next.next.data == 30 {
		fmt.Println("Linked list is OK")
	} else {
		fmt.Println("Linked list is incorrect!")
	}
}
