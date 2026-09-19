package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{
			input:  "   Hello world  ",
			output: []string{"hello", "world"},
		},
		{
			input:  "hello WORLD",
			output: []string{"hello", "world"},
		},
		{
			input:  "hElLo                 WoRlD",
			output: []string{"hello", "world"},
		},
		{
			input:  "hello WoRlD",
			output: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.output) {
			t.Errorf("Actual result %v is not of the required length: %d", actual, len(c.output))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.output[i]
			if word != expectedWord {
				t.Errorf("Actual word %s differs from expected word: %s", word, expectedWord)
				t.Fail()
			}
		}
	}
}
