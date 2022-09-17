package main

import (
	"testing"
)

func newBF() BloomFilter {
	var bf BloomFilter
	bf.filter_len = 32
	bf.barray = make([]uint8, bf.filter_len)
	return bf
}

func TestHash1(t *testing.T) {
	bf := newBF()
	t.Run("String is not empty", func(t *testing.T) {
		str := "0123456789"
		got := bf.Hash1(str)
		want := 13
		if got != want {
			t.Errorf("got: %d, want: %d, given %v", got, want, str)
		}
	})
	t.Run("String is empty", func(t *testing.T) {
		str := ""
		got := bf.Hash1(str)
		want := 0
		if got != want {
			t.Errorf("got: %d, want: %d, given %v", got, want, str)
		}
	})
}

func TestHash2(t *testing.T) {
	bf := newBF()
	t.Run("String is not empty", func(t *testing.T) {
		str := "0123456789"
		got := bf.Hash2(str)
		want := 5
		if got != want {
			t.Errorf("got: %d, want: %d, given %v", got, want, str)
		}
	})
	t.Run("String is empty", func(t *testing.T) {
		str := ""
		got := bf.Hash2(str)
		want := 0
		if got != want {
			t.Errorf("got: %d, want: %d, given %v", got, want, str)
		}
	})
}

func TestAdd(t *testing.T) {
	bf := newBF()
	t.Run("Add string", func(t *testing.T) {
		str := "0123456789"
		bf.Add(str)
		got := bf.barray[5]&1 != 0 && bf.barray[13]&1 != 0
		want := true
		if got != want {
			t.Errorf("got: %v, want: %v given %v", got, want, str)
		}
	})
}

func TestIsValue(t *testing.T) {
	t.Run("Filter is empty", func(t *testing.T) {
		str := "0123456789"
		bf := BloomFilter{filter_len: 32}
		got := bf.IsValue(str)
		want := false
		if got != want {
			t.Errorf("got: %v, want: %v given %v", got, want, str)
		}
	})
	bf := newBF()
	t.Run("String is in filter", func(t *testing.T) {
		str := "1234567890"
		bf.Add(str)
		got := bf.IsValue(str)
		want := true
		if got != want {
			t.Errorf("got: %v, want: %v given %v", got, want, str)
		}
	})

	t.Run("String is in filter", func(t *testing.T) {
		str := "2345678901"
		bf.Add(str)
		got := bf.IsValue(str)
		want := true
		if got != want {
			t.Errorf("got: %v, want: %v given %v", got, want, str)
		}
	})

	t.Run("String is not in filter (false positive)", func(t *testing.T) {
		str := "3456789012"
		got := bf.IsValue(str)
		want := true
		if got != want {
			t.Errorf("got: %v, want: %v given %v", got, want, str)
		}
	})
}
