package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{ // Test Case 1
			input: "	hello world		",
			expected: []string{"hello", "world"},
		},
		{ // Test Case 2
			input: "Krabby Arcanine",
			expected: []string{"krabby", "arcanine"},
		},
		{ // Test Case 3
			input: "  Good   day",
			expected: []string{"good", "day"},
		},
		{ // Test Case 4 - Empty input
			input: "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {											// Check the length of the returned slice
			t.Errorf("Length of actual result doesn't match the expected length")	// If it's different than expected, throw an error and continue
			continue
		}
		for i := range actual {														// Iterate over all words in returned slice and compare them to expected ones
			if actual[i] != c.expected[i] {
				t.Errorf("Words at index %d don't match", i)						// If words don't match, throw an error
			}
		}
	}
}
