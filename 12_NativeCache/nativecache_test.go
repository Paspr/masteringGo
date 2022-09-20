package main

import (
	"testing"
)

func TestInit(t *testing.T) {
	t.Run("Initialize new cache", func(t *testing.T) {
		nc := Init[string](5)

		got := [4]int{nc.size, len(nc.slots), len(nc.values), len(nc.hits)}
		want := [4]int{5, 5, 5, 5}

		if got != want {
			t.Errorf("got: %d, want: %d", got, want)
		}
	})
}

func TestHash(t *testing.T) {
	t.Run("String is not empty", func(t *testing.T) {
		nc := Init[string](5)
		str := "key1"

		got := nc.HashFun(str)
		want := 3

		if got != want {
			t.Errorf("got: %d, want: %d, given %v", got, want, str)
		}
	})
}

func TestSeekSlot(t *testing.T) {
	t.Run("There is free slot in cache", func(t *testing.T) {
		nc := Init[string](5)
		str := "key1"

		got := nc.SeekSlot(str)
		want := 3

		if got != want {
			t.Errorf("got: %d, want: %d, given %v", got, want, str)
		}
	})

	t.Run("There is no free slot in cache", func(t *testing.T) {
		nc := Init[string](1)
		str := "key1"
		val := "val for key1"
		nc.Put(str, val)
		str2 := "key2"

		got := nc.SeekSlot(str2)
		want := -1

		if got != want {
			t.Errorf("got: %d, want: %d, given %v", got, want, str)
		}
	})
}

func TestFindMin(t *testing.T) {
	t.Run("Find the least requested element's index", func(t *testing.T) {
		nc := Init[string](2)
		str := "key1"
		val := "val for key1"
		str2 := "key2"
		val2 := "val for key2"
		nc.Put(str, val)
		nc.Put(str2, val2)
		for i := 0; i <= 3; i++ {
			nc.Find(str)
		}
		for i := 0; i <= 2; i++ {
			nc.Find(str2)
		}

		got := nc.FindMin(nc.hits)
		want := 1

		if got != want {
			t.Errorf("got: %d, want: %d", got, want)
		}
	})
}

func TestFind(t *testing.T) {
	t.Run("Value is in cache", func(t *testing.T) {
		nc := Init[string](5)
		str := "key2"
		val := "val for key2"
		nc.Put(str, val)

		got, err := nc.Find(str)
		want := "val for key2"

		if got != want {
			t.Errorf("got: %v, want: %v, given %v", got, want, str)
		}

		if err != nil {
			t.Errorf("got error: %v, want: no error, given %v", err, str)
		}
	})

	t.Run("Value is not in cache", func(t *testing.T) {
		nc := Init[string](1)
		str := "key1"

		got, err := nc.Find(str)
		want := ""

		if err == nil {
			t.Errorf("expect error in case of missed value in cache")
		}

		if got != want {
			t.Errorf("got: %v, want: %v, given %v", got, want, str)
		}
	})
}

func TestPut(t *testing.T) {
	t.Run("Put value without replacing existing values", func(t *testing.T) {
		nc := Init[string](2)
		str := "key2"
		val := "val for key2"
		index := nc.SeekSlot("key2")
		nc.Put(str, val)

		indexGot := index != -1
		indexWant := true

		slotGot := nc.slots[index]
		slotWant := str

		valueGot := nc.values[index]
		valueWant := val

		if indexGot != indexWant {
			t.Errorf("got: %v, want: %v", indexGot, indexWant)
		}

		if slotGot != slotWant {
			t.Errorf("got: %v, want: %v", slotGot, slotWant)
		}

		if valueGot != valueWant {
			t.Errorf("got: %v, want: %v", valueGot, valueWant)
		}
	})

	t.Run("Put value with replacing existing values", func(t *testing.T) {
		nc := Init[string](2)
		str := "key1"
		val := "val for key1"
		str2 := "key2"
		val2 := "val for key2"
		nc.Put(str, val)
		nc.Put(str2, val2)
		nc.Find(str)
		str3 := "key3"
		val3 := "val for key3"
		nc.Put(str3, val3)

		got, err := nc.Find(str2)
		want := ""

		if err == nil {
			t.Errorf("expect error in case of missed value in cache")
		}

		if got != want {
			t.Errorf("got: %v, want: %v, given %v", got, want, str2)
		}

	})
}
