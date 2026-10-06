package main

import (
	//	"fmt"
	"testing"
	//	"cmp"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"whitespace-both-ends": {input: "  hello  world  ", expected: []string{"hello", "world"}},
		"tab-separator":        {input: "	hello		world	", expected: []string{"hello", "world"}},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("Assert slice length equal\nExpected:\t%d\nActual:\t%d", len(c.expected), len(actual))
				return
			}
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				if word != expectedWord {
					t.Errorf("Assert word equal\nExpected:\t%s\nActual:\t%s\n", expectedWord, word)
				}
			}
		})
	}
}
