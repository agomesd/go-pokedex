package utils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTryCatchPokemon(t *testing.T) {
	tests := map[string]struct {
		roll     float64
		baseExp  int
		expected bool
	}{
		"catch":  {roll: 0.10, baseExp: 123, expected: true},
		"escape": {roll: 0.9, baseExp: 299, expected: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := TryCatchPokemon(tc.roll, tc.baseExp)
			diff := cmp.Diff(tc.expected, got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
