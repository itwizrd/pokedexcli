package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "ThisIsOneWord",
			expected: []string{"thisisoneword"},
		},
		{
			input: "ALL UPPER CASE",
			expected: []string{"all","upper","case"}, 
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander","bulbasaur","pikachu"},
		},
		{
			input: "   Hello   THERE\tFriend  ",
			expected: []string{"hello","there","friend"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		lenA := len(actual)
		lenE := len(c.expected)
		if lenA != lenE {
			t.Errorf("length not expected, want %d; got %d", lenE, lenA)
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("wanted %s; got %s", expectedWord, word)
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
