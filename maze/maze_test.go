package main

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		inputfile string
		expected  int
	}{
		{
			"testfile1",
			4,
		},
		{
			"testfile2",
			-1,
		},
	}

	for _, test := range tests {
		result := resolver(test.inputfile)
		if result != test.expected {
			t.Errorf("can't resolve: %v, got: %v, expected: %v\n", test.inputfile, result, test.expected)
		}
	}
}
