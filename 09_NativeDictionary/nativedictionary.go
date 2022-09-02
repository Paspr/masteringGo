package main

import (
	"fmt"
	"os"
	"strconv"
)

type NativeDictionary[T any] struct {
	size   int
	slots  []string
	values []T
}

// создание экземпляра словаря
func Init[T any](sz int) NativeDictionary[T] {
	nd := NativeDictionary[T]{size: sz, slots: nil, values: nil}
	nd.slots = make([]string, sz)
	nd.values = make([]T, sz)
	return nd
}

func (nd *NativeDictionary[T]) HashFun(value string) int {
	// всегда возвращает корректный индекс слота
	barray := []byte(value)
	sum := 0
	for _, v := range barray {
		sum += int(v)
	}
	return sum % nd.size
}

func (nd *NativeDictionary[T]) IsKey(key string) bool {
	// возвращает true если ключ имеется
	index := nd.HashFun(key)
	return nd.slots[index] == key
}

func (nd *NativeDictionary[T]) Get(key string) (T, error) {
	// возвращает value для key,
	// или error если ключ не найден

	var result T

	index := nd.HashFun(key)
	if nd.slots[index] == key {
		result = nd.values[index]
		return result, nil
	}

	return result, fmt.Errorf("key '%s' is not found", key)
}

func (nd *NativeDictionary[T]) Put(key string, value T) {
	// гарантированно записываем
	// значение value по ключу key
	index := nd.HashFun(key)
	nd.slots[index] = key
	nd.values[index] = value
}
