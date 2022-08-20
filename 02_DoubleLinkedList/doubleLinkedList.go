package main

import (
	"errors"
	"os"
	"reflect"
)

type Node struct {
	prev  *Node
	next  *Node
	value int
}

type LinkedList2 struct {
	head *Node
	tail *Node
}

func (l *LinkedList2) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
		l.head.next = nil
		l.head.prev = nil
	} else {
		l.tail.next = &item
		item.prev = l.tail
	}

	l.tail = &item
	l.tail.next = nil
}

func (l *LinkedList2) Count() int {
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
func (l *LinkedList2) Find(n int) (Node, error) {
	currentNode := l.head
	for currentNode != nil {
		if currentNode.value == n {
			return *currentNode, nil
		}
		currentNode = currentNode.next
	}
	return Node{value: -1, next: nil}, errors.New("Node not found")
}

func (l *LinkedList2) FindAll(n int) []Node {
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

func (l *LinkedList2) Delete(n int, all bool) {
	deleteNode, _ := l.Find(n)
	if l.head == nil || deleteNode.value == -1 {
		return
	} else {
		temp := l.head
		for temp != nil {
			if temp.value == n {
				if temp == l.head {
					l.head = l.head.next
					if l.head == nil {
						l.tail = nil
						if !all {
							return
						}
					} else {
						l.head.prev = nil
						if !all {
							return
						}
					}
				} else {
					temp.prev.next = temp.next
					if temp.next == nil {
						l.tail = temp.prev
						if !all {
							return
						}
					} else {
						temp.next.prev = temp.prev
						if !all {
							return
						}
					}
				}
			}
			temp = temp.next
		}
	}
}

func (l *LinkedList2) Insert(after *Node, add Node) {
	currentNode := l.head
	if l.head == nil {
		l.InsertFirst(add)
	}
	for currentNode != nil {
		if currentNode.value == after.value {
			after.next = currentNode.next
			add.next = after.next
			add.prev = currentNode
			currentNode.next = &add
		}
		currentNode = currentNode.next
	}

}

func (l *LinkedList2) InsertFirst(first Node) {
	first.next = l.head
	first.prev = nil
	if l.head != nil {
		l.head.prev = &first
	}
	l.head = &first

}

func (l *LinkedList2) Clean() {
	l.head = nil
	l.tail = nil
}
