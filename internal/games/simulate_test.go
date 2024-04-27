package games

import "testing"

func TestGetStrategyFromNum(t *testing.T) {
	tc := []struct {
		name     string
		num      Weight
		weights  Weights
		expected int
	}{
		{
			name:     "Should get strategy correctly",
			weights:  Weights{0.1, 0, 0.2, 0.3, 0, 0.2, 0.1},
			num:      0.6,
			expected: 3,
		},
		{
			name:     "Should get strategy correctly",
			weights:  Weights{0.1, 0, 0.2, 0.3, 0, 0.2, 0.1},
			num:      0.4,
			expected: 3,
		},
		{
			name:     "Should get strategy correctly",
			weights:  Weights{0.5, 0.5},
			num:      0.5,
			expected: 0,
		},
		{
			name:     "Should get strategy correctly",
			weights:  Weights{0.5, 0.5},
			num:      0,
			expected: 0,
		},
		{
			name:     "Should get strategy correctly",
			weights:  Weights{0.5, 0.5},
			num:      0.6,
			expected: 1,
		},
		{
			name:     "Should get strategy correctly",
			weights:  Weights{0, 1},
			num:      0.6,
			expected: 1,
		},
		{
			name:     "Should get strategy correctly",
			weights:  Weights{0, 1},
			num:      0,
			expected: 1,
		},
		{
			name:     "Should get strategy correctly",
			weights:  Weights{0, 1},
			num:      1,
			expected: 1,
		},
		{
			name:     "Should get strategy correctly",
			weights:  Weights{0.7, 0.3},
			num:      0.6,
			expected: 0,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStrategyFromNum(tt.weights, tt.num); got != tt.expected {
				t.Fatalf(
					"Expected to get: %d, got: %d, val: %v, num: %v",
					tt.expected,
					got,
					tt.weights,
					tt.num,
				)
			}
		})
	}
}
