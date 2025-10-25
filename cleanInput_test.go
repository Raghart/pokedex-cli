package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello   world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "EEVEE SNORLAX CHARIZARD",
			expected: []string{"eevee", "snorlax", "charizard"},
		},
		{
			input:    "veNUSAUR UmbReOn TortErRa",
			expected: []string{"venusaur", "umbreon", "torterra"},
		},
	}
	for _, test := range cases {
		resultSlice := cleanInput(test.input)
		if len(resultSlice) != len(test.expected) {
			t.Errorf("the lengths of the expected array don't match with the arr: resultSlice=%v, expected=%d", len(resultSlice), len(test.expected))
		}
		for i := range resultSlice {
			word := resultSlice[i]
			expected := test.expected[i]

			if word != expected {
				t.Errorf("the expected words doesn't match: %v != %s", word, expected)
			}
		}
	}
}
