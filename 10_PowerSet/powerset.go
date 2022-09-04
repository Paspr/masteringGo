package main

import (
	"constraints"
	"os"
	"strconv"
)

type PowerSet[T constraints.Ordered] struct {
	// ваша реализация хранилища
	storage map[T]bool
}

// создание экземпляра множества
func Init[T constraints.Ordered]() PowerSet[T] {
	ps := PowerSet[T]{storage: nil}
	ps.storage = make(map[T]bool)
	return ps
}

func (ps *PowerSet[T]) Size() int {
	// количество элементов в множестве
	return len(ps.storage)
}

func (ps *PowerSet[T]) Put(value T) {
	// всегда срабатывает
	ps.storage[value] = true
}

func (ps *PowerSet[T]) Get(value T) bool {
	// возвращает true если value имеется в множестве
	return ps.storage[value]
}

func (ps *PowerSet[T]) Remove(value T) bool {
	// возвращает true если value удалено
	if ps.Get(value) {
		delete(ps.storage, value)
		return true
	}
	return false
}

func (ps *PowerSet[T]) Intersection(set2 PowerSet[T]) PowerSet[T] {
	// пересечение текущего множества и set2
	result := Init[T]()
	for v := range ps.storage {
		if set2.Get(v) {
			result.Put(v)
		}
	}
	return result
}

func (ps *PowerSet[T]) Union(set2 PowerSet[T]) PowerSet[T] {
	// объединение текущего множества и set2
	result := Init[T]()
	for v := range ps.storage {
		result.Put(v)
	}
	for v := range set2.storage {
		result.Put(v)
	}
	return result
}

func (ps *PowerSet[T]) Difference(set2 PowerSet[T]) PowerSet[T] {
	// разница текущего множества и set2
	result := Init[T]()
	for v := range ps.storage {
		if !set2.Get(v) {
			result.Put(v)
		}
	}
	return result
}

func (ps *PowerSet[T]) IsSubset(set2 PowerSet[T]) bool {
	// возвращает true, если set2 есть
	// подмножество текущего множества
	IsSubset := true
	for v := range set2.storage {
		if !ps.Get(v) {
			IsSubset = false
		}
	}
	return IsSubset
}
