package main


import (
	"os"
)

// битовый массив длиной f_len ...
type BloomFilter struct {
	filter_len int
	barray     []uint8
}

// хэш-функции
// 17
func (bf *BloomFilter) Hash1(s string) int {
	sum := 0
	for _, char := range s {
		code := int(char)
		sum = (sum*17 + code) % 32
	}
	return sum
}

// 223
func (bf *BloomFilter) Hash2(s string) int {
	sum := 0
	for _, char := range s {
		code := int(char)
		sum = (sum*223 + code) % 32
	}
	return sum
}

// добавляем строку s в фильтр
func (bf *BloomFilter) Add(s string) {
	if len(bf.barray) == 0 {
		bf.barray = make([]uint8, bf.filter_len)
	}
	index1 := bf.Hash1(s)
	index2 := bf.Hash2(s)
	bf.barray[index1] |= 1
	bf.barray[index2] |= 1

}

// проверка, имеется ли строка s в фильтре
func (bf *BloomFilter) IsValue(s string) bool {
	if len(bf.barray) == 0 {
		return false
	}
	index1 := bf.Hash1(s)
	index2 := bf.Hash2(s)
	return bf.barray[index1]&1 != 0 && bf.barray[index2]&1 != 0

}
