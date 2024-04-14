package games

import (
	"reflect"
	"slices"
	"testing"

	"github.com/hrvadl/algo/internal/matrix"
)

func TestCompleteMatrixToCompatible(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		expected *matrix.Matrix
	}{
		{
			name:     "Should handle empty matrix correctly",
			m:        matrix.Matrix{},
			expected: nil,
		},
		{
			name: "Should complete matrix correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{1, 1, 3},
					{2, -1, 3},
					{10, -1, 1},
				},
			},
			expected: &matrix.Matrix{
				Rows: []matrix.Row{
					{1, 1, 3, 1},
					{2, -1, 3, 1},
					{10, -1, 1, 1},
					{-1, -1, -1, 0},
				},
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := CompleteMatrixToCompatible(tt.m)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf("Expected to get: %v\ngot: %v", tt.expected, got)
			}
		})
	}
}

func TestGetGameWeight(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		expected float64
	}{
		{
			name: "Should find game weight correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{0.33, 0, 1.33, 1.33, 0.33},
					{1, -1.5, 4, 7.5, 0.5},
					{-0.33, 0.5, -0.33, 0.17, 0.17},
					{0, 0.5, 0, 0.5, 0.5},
				},
			},
			expected: 2,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := GetGameWeight(tt.m); got != tt.expected {
				t.Fatalf("Expected to get: %v, got: %v", tt.expected, got)
			}
		})
	}
}

func TestCorrectGameWeight(t *testing.T) {
	tc := []struct {
		name     string
		w        float64
		minabs   float64
		expected float64
	}{
		{
			name:     "Should find game weight correctly",
			w:        2,
			minabs:   1,
			expected: 1,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := CorrectGameWeight(tt.w, tt.minabs); got != tt.expected {
				t.Fatalf("Expected to get: %v, got: %v", tt.expected, got)
			}
		})
	}
}

func TestGetCorrectedGameWeight(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		minabs   float64
		expected float64
	}{
		{
			name: "Should find corrected game weight correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{0.33, 0, 1.33, 1.33, 0.33},
					{1, -1.5, 4, 7.5, 0.5},
					{-0.33, 0.5, -0.33, 0.17, 0.17},
					{0, 0.5, 0, 0.5, 0.5},
				},
			},
			minabs:   1,
			expected: 1,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := CorrectGameWeight(GetGameWeight(tt.m), tt.minabs); got != tt.expected {
				t.Fatalf("Expected to get: %v, got: %v", tt.expected, got)
			}
		})
	}
}

func TestCorrectMixedStrategy(t *testing.T) {
	tc := []struct {
		name     string
		strategy []float64
		minabs   float64
		expected []float64
	}{
		{
			name:     "Should find corrected game weight correctly",
			minabs:   2,
			strategy: []float64{1, 2, 3},
			expected: []float64{2, 4, 6},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := CorrectMixedStrategy(tt.strategy, tt.minabs); !slices.Equal(
				got,
				tt.expected,
			) {
				t.Fatalf("Expected to get: %v, got: %v", tt.expected, got)
			}
		})
	}
}
