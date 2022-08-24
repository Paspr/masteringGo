package main

import (
	"os"
	"fmt"
)

type DynArray[T any] struct {
	count    int
	capacity int
	array    []T
}

func (da *DynArray[T]) Init() {
	da.count = 0
	da.MakeArray(16)
}

func (da *DynArray[T]) MakeArray(sz int) {
	var arr = make([]T, sz)
	copy(arr, da.array) //  копируем содержимое array в arr ...
	da.capacity = sz
	da.array = arr //
}

func (da *DynArray[T]) Insert(itm T, index int) error {
	if index == da.count {
		da.Append(itm)
		return nil
	}
	if index > da.count || index < 0 {
		return fmt.Errorf("bad index '%d'", index)
	} else if da.count+1 > da.capacity {
		da.MakeArray(da.capacity * 2)
	}

	tempRange := da.count - index
	tempArr := da.array[index : index+tempRange]
	var arr = make([]T, da.capacity)
	copy(arr, tempArr)
	for i, j := index+1, 0; j < tempRange; i, j = i+1, j+1 {
		da.array[i] = arr[j]
	}

	da.array[index] = itm
	da.count++
	return nil
}

func (da *DynArray[T]) Remove(index int) error {
	if index >= da.count || index < 0 {
		return fmt.Errorf("bad index '%d'", index)
	}
	for i := index; i < da.capacity-1; i++ {
		da.array[i] = da.array[i+1]
	}
	da.count--
	if da.count < da.capacity/2 {
		if da.count < 16 {
			da.MakeArray(16)
		} else {
			da.MakeArray(int(float64(da.capacity) / 1.5))
		}
	}
	return nil
}

func (da *DynArray[T]) Append(itm T) {
	if da.count+1 > da.capacity {
		da.MakeArray(da.capacity * 2)
	}
	da.array[da.count] = itm
	da.count++
}

func (da *DynArray[T]) GetItem(index int) (T, error) {
	var result T
	if index >= da.count || index < 0 {
		return result, fmt.Errorf("bad index '%d'", index)
	} else {
		result = da.array[index]
	}
	return result, nil
}
