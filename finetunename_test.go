package main

import (
	
	"testing"
)

func TestFineTuneName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Mr Emeka Aneke", "Mr Emeka"},
		{"Aunty Gloria HT", "Mrs Gloria"},
		{"Uncle Emeka Aneke", "Mr Emeka"},
		{"Frasco Law", "Frasco"},
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := fineTuneName(test.input)
		if result != test.expected {
			t.Errorf("Expected %s for input %s, but got %s", test.expected, test.input, result)
		}
	}
}
