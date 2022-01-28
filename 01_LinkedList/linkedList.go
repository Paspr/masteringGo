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

func (n LinkedList) FindAll(data int) []Node {
	// Finds all nodes by the specified data and returns a slice of nodes
	var nodes []Node
	currentNode := n.head
	for currentNode != nil {
		if currentNode.data == data {
			nodes = append(nodes, *currentNode)
		}
		currentNode = currentNode.next
	}
	return nodes
}

// Remove (bool)
func (n *LinkedList) Remove(data int) bool {
	if n.Find(data) == nil || n.head == nil {
		return false
	} else {
		currentNode := n.head
		var previousNode *Node
		for currentNode != nil {
			if currentNode.data == data {
				if n.head.next == nil {
					n.head = nil
					n.tail = nil
					return true
				}
				if n.head.data == currentNode.data {
					n.head = n.head.next
					return true
				} else {
					if currentNode.next == nil {
						previousNode.next = nil
						n.tail = previousNode
						return true
					} else {
						previousNode.next = currentNode.next
						return true
					}
				}

			}
			previousNode = currentNode
			currentNode = currentNode.next
		}
		return true
	}

}

// RemoveAll
// InsertAfter

func main() {
	var testList LinkedList
	testNode := Node{data: 10}
	testNode2 := Node{data: 20}
	testNode3 := Node{data: 30}
	testNode4 := Node{data: 30}

	testList.addInTail(&testNode)
	testList.addInTail(&testNode2)
	testList.addInTail(&testNode3)
	testList.addInTail(&testNode4)

	fmt.Println("initial list:")
	testList.printLinkedList()
	fmt.Println(testList.Count())
	fmt.Println("FindAll result:")
	fmt.Println(testList.FindAll(30))
	fmt.Println("Remove '30' value")
	fmt.Println(testList.Remove(30))
	fmt.Println("list after remove:")
	testList.printLinkedList()
	fmt.Println(testList.Count())

	if testList.head.data == 10 && testList.head.next.data == 20 && testList.head.next.next.data == 30 {
		fmt.Println("Linked list is OK")
	} else {
		fmt.Println("Linked list is incorrect!")
	}

}
