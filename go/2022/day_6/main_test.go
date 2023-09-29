package main

import "testing"

func TestIsStringSliceUnique(t *testing.T) {
	cases := []struct {
		input  string
		result bool
	}{
		{input: "aa",
			result: false},
		{input: "a",
			result: true},
	}

	for _, c := range cases {
		result := IsStringSliceUnique(c.input)

		if result != c.result {
			t.Logf("expected %t but got: %t", c.result, result)
			t.Fail()
		}
	}
}
