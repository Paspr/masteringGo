package main

import (
	"os"
	"fmt"
)

type Queue[T any] struct {
	array []T
}

func (q *Queue[T]) Size() int {
	return len(q.array)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var result T
	queueSize := q.Size()
	if queueSize == 0 {
		return result, fmt.Errorf("queue size is '%d'", queueSize)
	}
	result = q.array[0]
	q.array = q.array[1:]
	return result, nil
}

func (q *Queue[T]) Enqueue(itm T) {
	q.array = append(q.array, itm)
}
