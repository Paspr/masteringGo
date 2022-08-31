package main

import (
	"constraints"
	"errors"
	"os"
)

type Node[T constraints.Ordered] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type OrderedList[T constraints.Ordered] struct {
	head       *Node[T]
	tail       *Node[T]
	_ascending bool
}

func (l *OrderedList[T]) Count() int {
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

func (l *OrderedList[T]) Add(item T) {
	var temp Node[T]
	temp.value = item
	if l.head == nil {
		temp.next = l.head
		l.head = &temp
		l.tail = l.head
	} else {
		currentNode := l.head
		if l._ascending {
			for currentNode != nil && l.Compare(currentNode.value, temp.value) == -1 {
				currentNode = currentNode.next
			}
		} else {
			for currentNode != nil && l.Compare(currentNode.value, temp.value) == 1 {
				currentNode = currentNode.next
			}
		}
		if currentNode == l.head {
			currentNode.prev = &temp
			currentNode.prev.next = l.head
			l.head = currentNode.prev
		} else {
			if currentNode != nil {
				currentNode.prev.next = &temp
				currentNode.prev.next.next = currentNode
				currentNode.prev.next.prev = currentNode.prev
				currentNode.prev = currentNode.prev.next
			} else {
				l.tail.next = &temp
				l.tail.next.prev = l.tail
				l.tail = l.tail.next
			}
		}
	}
}

func (l *OrderedList[T]) Find(n T) (Node[T], error) {
	currentNode := l.head
	if l._ascending {
		for currentNode != nil &&
			(l.Compare(currentNode.value, n) == -1 || l.Compare(currentNode.value, n) == 0) {
			if l.Compare(currentNode.value, n) == 0 {
				return *currentNode, nil
			}
			currentNode = currentNode.next
		}
	} else {
		for currentNode != nil &&
			(l.Compare(currentNode.value, n) == 1 || l.Compare(currentNode.value, n) == 0) {
			if l.Compare(currentNode.value, n) == 0 {
				return *currentNode, nil
			}
			currentNode = currentNode.next
		}
	}
	return Node[T]{value: n, next: nil, prev: nil}, errors.New("Node not found")
}

func (l *OrderedList[T]) Delete(n T) {
	_, err := l.Find(n)
	if l.head == nil || err != nil {
		return
	} else {
		temp := l.head
		for temp != nil {
			if l.Compare(temp.value, n) == 0 {
				if temp == l.head {
					l.head = l.head.next
					if l.head == nil {
						l.tail = nil
						return
					} else {
						l.head.prev = nil
						return
					}
				} else {
					temp.prev.next = temp.next
					if temp.next == nil {
						l.tail = temp.prev
						return
					} else {
						temp.next.prev = temp.prev
						return
					}
				}
			}
			temp = temp.next
		}
	}
}

func (l *OrderedList[T]) Clear(asc bool) {
	l._ascending = asc
	l.head = nil
	l.tail = nil

}

func (l *OrderedList[T]) Compare(v1 T, v2 T) int {
	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return +1
	}
	return 0
}
