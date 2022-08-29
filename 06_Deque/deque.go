package main

import (
	"os"
	"fmt"
)

type Deque[T any] struct {
	array []T
}

func (d *Deque[T]) Size() int {
	return len(d.array)
}

func (d *Deque[T]) AddFront(itm T) {
	temp := make([]T, d.Size()+1)
	temp[0] = itm
	copy(temp[1:], d.array)
	d.array = temp
}

func (d *Deque[T]) AddTail(itm T) {
	d.array = append(d.array, itm)
}

func (d *Deque[T]) RemoveFront() (T, error) {
	var result T
	dequeSize := d.Size()
	if dequeSize == 0 {
		return result, fmt.Errorf("deque size is '%d'", dequeSize)
	}
	result = d.array[0]
	d.array = d.array[1:]
	return result, nil
}

func (d *Deque[T]) RemoveTail() (T, error) {
	var result T
	dequeSize := d.Size()
	if dequeSize == 0 {
		return result, fmt.Errorf("deque size is '%d'", dequeSize)
	}
	result = d.array[dequeSize-1]
	d.array = d.array[:dequeSize-1]
	return result, nil
}
