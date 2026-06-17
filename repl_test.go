package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    // ...
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
	// add more cases here
	{
		input: "  HELLO WORLD  ",
		expected: []string{"hello", "world"},
	},
	{
		input: "  hello   world  ",
		expected: []string{"hello", "world"},
	},
	{
		input: "  hello   world   ",
		expected: []string{"hello", "world"},
	},
	{
		input: "  hello   world   ",
		expected: []string{"hello", "world"},
	},
	}

	// then loop over the cases and run the tests

	for _, c := range cases {
		actual := cleanInput(c.input)
		// check the length of the actual slice against the expected slice if they don't match, use t.Errorf to print an error message and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) = %v; expected %v", c.input, actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// check each word in the slice, if they don't match, use t.Errorf to print an error message and fail the test
			if word != expectedWord {
				t.Errorf("cleanInput(%q) = %v; expected %v", c.input, actual, c.expected)
			}

		}
	}
}