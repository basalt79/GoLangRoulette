package util

import (
	"testing"
)

func TestInBetween(t *testing.T) {
	tests := []struct {
		in  int
		min int
		max int
		out bool
	}{
		{2, 0, 3, true},
		{44, 0, 3, false},
		{22, 0, 36, true},
		{1, 1, 2, true},
		{2, 1, 2, true},
	}

	for _, test := range tests {
		result := InBetween(test.in, test.min, test.max)
		if result != test.out {
			t.Errorf("number %v bewteen %v and %v returned %v", test.in, test.min, test.max, test.out)
		}
	}
}
