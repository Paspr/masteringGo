package main


import (
	"os"
)

// битовый массив длиной f_len ...
type BloomFilter struct {
	filter_len int
	barray     []bool
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
	index1 := bf.Hash1(s)
	index2 := bf.Hash2(s)
	bf.barray[index1] = true
	bf.barray[index2] = true

}

// проверка, имеется ли строка s в фильтре
func (bf *BloomFilter) IsValue(s string) bool {
	index1 := bf.Hash1(s)
	index2 := bf.Hash2(s)
	return bf.barray[index1] && bf.barray[index2]

}
