package sliceh

import (
	"slices"
	"testing"
)

func TestMinIdx(t *testing.T) {
	tc := []struct {
		name     string
		s        []int
		expected []int
	}{
		{
			name:     "should find min indexes correctly",
			s:        []int{1, 2, 3, 4, 1},
			expected: []int{0, 4},
		},
		{
			name:     "should find min indexes correctly",
			s:        []int{1, 2, -33, 4, 1},
			expected: []int{2},
		},
		{
			name:     "should find min indexes correctly",
			s:        []int{},
			expected: nil,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := MinIdxs(tt.s); !slices.Equal(got, tt.expected) {
				t.Fatalf("Expected to get: %v, got: %v", tt.expected, got)
			}
		})
	}
}

func TestMaxIdx(t *testing.T) {
	tc := []struct {
		name     string
		s        []int
		expected []int
	}{
		{
			name:     "should find max indexes correctly",
			s:        []int{4, 2, 3, 4, 1},
			expected: []int{0, 3},
		},
		{
			name:     "should find max indexes correctly",
			s:        []int{1, -1, -33, -2, 1},
			expected: []int{0, 4},
		},
		{
			name:     "should find max indexes correctly",
			s:        []int{},
			expected: nil,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := MaxIdxs(tt.s); !slices.Equal(got, tt.expected) {
				t.Fatalf("Expected to get: %v, got: %v", tt.expected, got)
			}
		})
	}
}
