package utils

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected []string
	}{
		"simple":        {input: "hello world", expected: []string{"hello", "world"}},
		"space padded":  {input: " hello world ", expected: []string{"hello", "world"}},
		"with capitals": {input: "Hello World", expected: []string{"hello", "world"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CleanInput(tc.input)
			diff := cmp.Diff(tc.expected, got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
