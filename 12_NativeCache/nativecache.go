package main

import (
	"fmt"
)

type NativeCache[T any] struct {
	size   int
	slots  []string
	values []T
	hits   []int
}

func Init[T any](sz int) NativeCache[T] {
	// initializing new cache
	nc := NativeCache[T]{size: sz, slots: nil, values: nil}
	nc.slots = make([]string, sz)
	nc.values = make([]T, sz)
	nc.hits = make([]int, sz)
	return nc
}

func (nc *NativeCache[T]) HashFun(value string) int {
	// return slot index
	barray := []byte(value)
	sum := 0
	for _, v := range barray {
		sum += int(v)
	}
	return sum % nc.size
}

func (nc *NativeCache[T]) SeekSlot(value string) int {
	// finds empty slot (if available),
	// otherwise returns -1
	index := nc.HashFun(value)
	if nc.slots[index] == "" {
		return index
	} else {
		var i = index
		for nc.slots[index] != "" {
			index = index + 1
			if index >= nc.size {
				index = index - nc.size
			}
			if nc.slots[index] == "" {
				return index
			}
			if nc.slots[index] != "" && index == i {
				break
			}
		}
	}
	return -1
}

func (nc *NativeCache[T]) FindMin(hits []int) int {
	// finds the least requested element's index
	min := nc.hits[0]
	minIndex := 0
	for i := 0; i < len(nc.hits); i++ {
		if min > nc.hits[i] {
			min = nc.hits[i]
			minIndex = i
		}
	}
	return minIndex
}

func (nc *NativeCache[T]) Put(key string, value T) {
	// put key-value pair into cache
	index := nc.SeekSlot(key)
	if index != -1 {
		nc.slots[index] = key
		nc.values[index] = value
	} else {
		nc.slots[nc.FindMin(nc.hits)] = key
		nc.values[nc.FindMin(nc.hits)] = value
		nc.hits[nc.FindMin(nc.hits)] = 0

	}
}

func (nc *NativeCache[T]) Find(value string) (T, error) {
	// searches for value in cache
	index := nc.HashFun(value)
	var result T
	if nc.slots[index] == value {
		nc.hits[index] = nc.hits[index] + 1
		result = nc.values[index]
		return result, nil
	} else {
		var i = index
		for nc.slots[index] != "" {
			index = index + 1
			if index >= nc.size {
				index = index - nc.size
			}
			if nc.slots[index] == "" {
				break
			} else {
				if nc.slots[index] == value {
					nc.hits[index] = nc.hits[index] + 1
					result = nc.values[index]
					return result, nil
				}
				if nc.slots[index] != "" && index == i {
					break
				}
			}
		}
	}
	return result, fmt.Errorf("value '%s' is not found", value)
}
