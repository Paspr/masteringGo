package main


import (
    "strconv"
	"os"
)


type HashTable struct {
	size  int
	step  int
	slots []string
}

func Init(sz int, stp int) HashTable {
	ht := HashTable{size: sz, step: stp, slots: nil}
	ht.slots = make([]string, sz)
	return ht
}

func (ht *HashTable) HashFun(value string) int {
	// всегда возвращает корректный индекс слота
	barray := []byte(value)
	sum := 0
	for _, v := range barray {
		sum += int(v)
	}
	return sum % ht.size
}

func (ht *HashTable) SeekSlot(value string) int {
	// находит индекс пустого слота для значения,
	// или -1
	index := ht.HashFun(value)
	if ht.slots[index] == "" {
		return index
	} else {
		var i = index
		for ht.slots[index] != "" {
			index = index + ht.step
			if index >= ht.size {
				index = index - ht.size
			}
			if ht.slots[index] == "" {
				return index
			}
			if ht.slots[index] != "" && index == i {
				break
			}
		}
	}
	return -1
}

func (ht *HashTable) Put(value string) int {
	// записываем значение по хэш-функции

	// возвращается индекс слота или -1
	// если из-за коллизий элемент не удаётся разместить
	index := ht.SeekSlot(value)
	if index != -1 {
		ht.slots[index] = value
		return index
	}
	return -1
}

func (ht *HashTable) Find(value string) int {
	// находит индекс слота со значением, или -1
	index := ht.HashFun(value)
	if ht.slots[index] == value {
		return index
	} else {
		var i = index
		for ht.slots[index] != "" {
			index = index + ht.step
			if index >= ht.size {
				index = index - ht.size
			}
			if ht.slots[index] == "" {
				break
			} else {
				if ht.slots[index] == value {
					return index
				}
				if ht.slots[index] != "" && index == i {
					break
				}
			}
		}
	}
	return -1
}
