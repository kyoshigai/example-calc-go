package calc

import "testing"

func TestSum(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "OK", input: []int{0, 1, 2}, expected: 3},
		{name: "includes negative", input: []int{0, 1, -4}, expected: -3},
		{name: "includes only zero", input: []int{0}, expected: 0},
		{name: "empty slice", input: []int{}, expected: 0},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.input)
			if got != tt.expected {
				t.Errorf("expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
