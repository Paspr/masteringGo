package main

import (
    "os"
    "reflect"
    "errors"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}

	l.tail = &item
}

func (l *LinkedList) Count() int {
	if l.head == nil {
		return 0
	} else {
		currentNode := l.head
		length := 0
		for currentNode != nil {
			length++
			currentNode = currentNode.next
		}
		return length
	}
}

// error не nil, если узел не найден
func (l *LinkedList) Find(n int) (Node, error) {
	currentNode := l.head
	for currentNode != nil {
		if currentNode.value == n {
			return *currentNode, nil
		}
		currentNode = currentNode.next
	}
	return Node{value: -1, next: nil}, errors.New("Node not found")
}

func (l *LinkedList) FindAll(n int) []Node {
	var nodes []Node
	currentNode := l.head
	for currentNode != nil {
		if currentNode.value == n {
			nodes = append(nodes, *currentNode)
		}
		currentNode = currentNode.next
	}
	return nodes
}

func (l *LinkedList) Delete(n int, all bool) {

	innerDelete := func(val int) {

		currentNode := l.head
		var previousNode *Node
		for currentNode != nil {
			if currentNode.value == val {
				if l.head.next == nil {
					l.head = nil
					l.tail = nil
					return
				}
				if l.head.value == currentNode.value {
					l.head = l.head.next
					return
				} else {
					if currentNode.next == nil {
						previousNode.next = nil
						l.tail = previousNode
						return
					} else {
						previousNode.next = currentNode.next
						return
					}
				}

			}
			previousNode = currentNode
			currentNode = currentNode.next
		}
		return
	}

	if l.head == nil {
		return
	} else if !all {
		innerDelete(n)
	} else {
		length := l.Count()
		for i := 0; i < length; i++ {
			innerDelete(n)

		}

	}
}

func (l *LinkedList) Insert(after *Node, add Node) {
	currentNode := l.head
	if after == nil && l.head == nil {
		add.next = l.head
		l.head = &add
		l.tail = l.head
	}
	for currentNode != nil {
		if currentNode.value == after.value {
			after.next = currentNode.next
			add.next = after.next
			currentNode.next = &add
			if l.tail.next != nil {
				l.tail = currentNode.next
			}
		}
		currentNode = currentNode.next
	}

}

func (l *LinkedList) InsertFirst(first Node) {
	first.next = l.head
	l.head = &first
	l.tail = l.head

}

func (l *LinkedList) Clean() {
	l.head = nil
	l.tail = nil
}
