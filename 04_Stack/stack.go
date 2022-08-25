package main

import (
	"os"
	"fmt"
)

type Stack[T any] struct {
	array []T
}

func (st *Stack[T]) Size() int {
	return len(st.array)
}

func (st *Stack[T]) Peek() (T, error) {
	var result T
	stackSize := st.Size()
	if stackSize == 0 {
		return result, fmt.Errorf("stack size is '%d'", stackSize)
	}
	result = st.array[stackSize-1]
	return result, nil
}

func (st *Stack[T]) Pop() (T, error) {
	var result T
	stackSize := st.Size()
	if stackSize == 0 {
		return result, fmt.Errorf("stack size is '%d'", stackSize)
	}
	result = st.array[stackSize-1]
	st.array = st.array[:stackSize-1]
	return result, nil
}

func (st *Stack[T]) Push(itm T) {
	st.array = append(st.array, itm)
}
