package main

import (
	"testing"
   "strings"
)

func Test(t *testing.T) {
	var tests = []struct {
		n        int32
		k        int32
		s        string
		expected string
	}{
		{
			4, 1, "3943", "3993",
		},
		{
			6, 3, "092282", "992299",
		},
		{
			4, 1, "0011", "-1",
		},
		{
			4, 4, "3943", "9999",
		},
		{
			5, 3, "39743", "99799",
		},
	}

	for _, test := range tests {
		r := highestValuePalindrome(test.s, test.n, test.k)

		if strings.Compare(r, test.expected) != 0 {
			t.Fatalf("s:%s, expected:%s, got:%s\n", test.s, test.expected, r)
		}
	}
}
