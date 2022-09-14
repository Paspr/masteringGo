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
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "String is not empty",
			input: "0123456789",
			want:  13,
		},
		{
			name:  "String is empty",
			input: "",
			want:  0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := bf.Hash1(testCase.input)
			if testCase.want != got {
				t.Errorf("got: %d, want: %d", got, testCase.want)
			}
		})
	}
}

func TestHash2(t *testing.T) {
	bf := newBF()
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "String is not empty",
			input: "0123456789",
			want:  5,
		},
		{
			name:  "String is empty",
			input: "",
			want:  0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := bf.Hash2(testCase.input)
			if testCase.want != got {
				t.Errorf("got: %d, want: %d", got, testCase.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	bf := newBF()
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "String is not empty",
			input: "0123456789",
			want:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			bf.Add(testCase.input)
			got := bf.barray[5]&1 != 0 && bf.barray[13]&1 != 0
			if testCase.want != got {
				t.Errorf("got: %v, want: %v", got, testCase.want)
			}
		})
	}
}

func TestIsValue(t *testing.T) {
	bf := newBF()
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "String is in filter",
			input: "0123456789",
			want:  true,
		},
		{
			name:  "String is in filter",
			input: "1234567890",
			want:  true,
		},
		{
			name:  "String is in filter",
			input: "2345678901",
			want:  true,
		},
		{
			name:  "String is not in filter (false positive)",
			input: "3456789012",
			want:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.input != "3456789012" {
				bf.Add(testCase.input)
			}
			got := bf.IsValue(testCase.input)
			if testCase.want != got {
				t.Errorf("got: %v, want: %v", got, testCase.want)
			}
		})
	}
}
