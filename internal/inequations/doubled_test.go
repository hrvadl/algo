package inequations

import (
	"reflect"
	"testing"

	"github.com/hrvadl/algo/internal/matrix"
)

func TestFindMaxDoubledWithSupportSolution(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		expected DoubledSolution
	}{
		{
			name: "Should calculate doubled support solution properly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{1, 1, -1, -2, 6},
					{-1, -1, -1, -1, -5},
					{2, -1, 3, 4, 10},
					{-1, -2, 1, 1, 0},
				},
				TopTitle: []matrix.Variable{
					{
						FirstStageName:   "x",
						FirstStageIndex:  0,
						SecondStageName:  "v",
						SecondStageIndex: 0,
					},
					{
						FirstStageName:   "x",
						FirstStageIndex:  1,
						SecondStageName:  "v",
						SecondStageIndex: 1,
					},
					{
						FirstStageName:   "x",
						FirstStageIndex:  2,
						SecondStageName:  "v",
						SecondStageIndex: 2,
					},
					{
						FirstStageName:   "x",
						FirstStageIndex:  3,
						SecondStageName:  "v",
						SecondStageIndex: 3,
					},
					{
						FirstStageName:  "w",
						SecondStageName: "1",
					},
				},
				LeftTitle: []matrix.Variable{
					{
						FirstStageName:   "y",
						FirstStageIndex:  0,
						SecondStageName:  "u",
						SecondStageIndex: 0,
					},
					{
						FirstStageName:   "y",
						FirstStageIndex:  1,
						SecondStageName:  "u",
						SecondStageIndex: 1,
					},
					{
						FirstStageName:   "y",
						FirstStageIndex:  2,
						SecondStageName:  "u",
						SecondStageIndex: 2,
					},
					{
						FirstStageName:  "z",
						SecondStageName: "1",
					},
				},
			},
			expected: DoubledSolution{
				Min: Solution{
					Result: []float64{0, -1, 0},
					Matrix: matrix.Matrix{
						Rows: []matrix.Row{
							{1, 0, -2, -3, 1},
							{-1, 1, 1, 1, 5},
							{2, -3, 1, 2, 0},
							{-1, -1, 2, 2, 5},
						},
						TopTitle: []matrix.Variable{
							{
								FirstStageName:   "y",
								FirstStageIndex:  1,
								SecondStageName:  "u",
								SecondStageIndex: 1,
							},
							{
								FirstStageName:   "x",
								FirstStageIndex:  1,
								SecondStageName:  "v",
								SecondStageIndex: 1,
							},
							{
								FirstStageName:   "x",
								FirstStageIndex:  2,
								SecondStageName:  "v",
								SecondStageIndex: 2,
							},
							{
								FirstStageName:   "x",
								FirstStageIndex:  3,
								SecondStageName:  "v",
								SecondStageIndex: 3,
							},
							{
								FirstStageName:  "w",
								SecondStageName: "1",
							},
						},
						LeftTitle: []matrix.Variable{
							{
								FirstStageName:   "y",
								FirstStageIndex:  0,
								SecondStageName:  "u",
								SecondStageIndex: 0,
							},
							{
								FirstStageName:   "x",
								FirstStageIndex:  0,
								SecondStageName:  "v",
								SecondStageIndex: 0,
							},
							{
								FirstStageName:   "y",
								FirstStageIndex:  2,
								SecondStageName:  "u",
								SecondStageIndex: 2,
							},
							{
								FirstStageName:  "z",
								SecondStageName: "1",
							},
						},
					},
				},
				Max: Solution{
					Result: []float64{5, 0, 0, 0},
					Matrix: matrix.Matrix{
						Rows: []matrix.Row{
							{1, 0, -2, -3, 1},
							{-1, 1, 1, 1, 5},
							{2, -3, 1, 2, 0},
							{-1, -1, 2, 2, 5},
						},
						TopTitle: []matrix.Variable{
							{
								FirstStageName:   "y",
								FirstStageIndex:  1,
								SecondStageName:  "u",
								SecondStageIndex: 1,
							},
							{
								FirstStageName:   "x",
								FirstStageIndex:  1,
								SecondStageName:  "v",
								SecondStageIndex: 1,
							},
							{
								FirstStageName:   "x",
								FirstStageIndex:  2,
								SecondStageName:  "v",
								SecondStageIndex: 2,
							},
							{
								FirstStageName:   "x",
								FirstStageIndex:  3,
								SecondStageName:  "v",
								SecondStageIndex: 3,
							},
							{
								FirstStageName:  "w",
								SecondStageName: "1",
							},
						},
						LeftTitle: []matrix.Variable{
							{
								FirstStageName:   "y",
								FirstStageIndex:  0,
								SecondStageName:  "u",
								SecondStageIndex: 0,
							},
							{
								FirstStageName:   "x",
								FirstStageIndex:  0,
								SecondStageName:  "v",
								SecondStageIndex: 0,
							},
							{
								FirstStageName:   "y",
								FirstStageIndex:  2,
								SecondStageName:  "u",
								SecondStageIndex: 2,
							},
							{
								FirstStageName:  "z",
								SecondStageName: "1",
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if actual, _ := FindSupportSolution(tt.m); !reflect.DeepEqual(
				actual,
				tt.expected.Max,
			) {
				t.Fatalf("want: %v\ngot:%v", tt.expected, actual)
			}
		})
	}
}

func TestFindMaxDoubledWithOptimalSolution(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		expected DoubledOptimalSolution
	}{
		{
			name: "Should calculate doubled support solution properly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{-2, 3, 14},
					{1, 1, 8},
					{-2, -7, 0},
				},
			},
			expected: DoubledOptimalSolution{
				MinSolution: MinSolution{
					Min: 46,
					Solution: Solution{
						Result: []float64{1, 4},
						Matrix: matrix.Matrix{
							LeftTitle: []matrix.Variable{
								{
									FirstStageName:   "x",
									FirstStageIndex:  1,
									SecondStageName:  "v",
									SecondStageIndex: 1,
								},
								{
									FirstStageName:   "x",
									FirstStageIndex:  0,
									SecondStageName:  "v",
									SecondStageIndex: 0,
								},
								{
									FirstStageName:  "z",
									SecondStageName: "1",
								},
							},
							TopTitle: []matrix.Variable{
								{
									FirstStageName:   "y",
									FirstStageIndex:  1,
									SecondStageName:  "u",
									SecondStageIndex: 1,
								},
								{
									FirstStageName:   "y",
									FirstStageIndex:  0,
									SecondStageName:  "u",
									SecondStageIndex: 0,
								},
								{
									FirstStageName:  "1",
									SecondStageName: "w",
								},
							},
							Rows: []matrix.Row{
								{0.4, 0.2, 6},
								{0.6, -0.2, 2},
								{4, 1, 46},
							},
						},
					},
				},
				MaxSolution: MaxSolution{
					Max: 46,
					Solution: Solution{
						Result: []float64{2, 6},
						Matrix: matrix.Matrix{
							Rows: []matrix.Row{
								{0.4, 0.2, 6},
								{0.6, -0.2, 2},
								{4, 1, 46},
							},
							LeftTitle: []matrix.Variable{
								{
									FirstStageName:   "x",
									FirstStageIndex:  1,
									SecondStageName:  "v",
									SecondStageIndex: 1,
								},
								{
									FirstStageName:   "x",
									FirstStageIndex:  0,
									SecondStageName:  "v",
									SecondStageIndex: 0,
								},
								{
									FirstStageName:  "z",
									SecondStageName: "1",
								},
							},
							TopTitle: []matrix.Variable{
								{
									FirstStageName:   "y",
									FirstStageIndex:  1,
									SecondStageName:  "u",
									SecondStageIndex: 1,
								},
								{
									FirstStageName:   "y",
									FirstStageIndex:  0,
									SecondStageName:  "u",
									SecondStageIndex: 0,
								},
								{
									FirstStageName:  "1",
									SecondStageName: "w",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Should calculate doubled support solution properly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{1, 1, -1, -2, 6},
					{-1, -1, -1, -1, -5},
					{2, -1, 3, 4, 10},
					{-1, -2, 1, 1, 0},
				},
				TopTitle: []matrix.Variable{
					{
						FirstStageName:   "x",
						FirstStageIndex:  0,
						SecondStageName:  "v",
						SecondStageIndex: 0,
					},
					{
						FirstStageName:   "x",
						FirstStageIndex:  1,
						SecondStageName:  "v",
						SecondStageIndex: 1,
					},
					{
						FirstStageName:   "x",
						FirstStageIndex:  2,
						SecondStageName:  "v",
						SecondStageIndex: 2,
					},
					{
						FirstStageName:   "x",
						FirstStageIndex:  3,
						SecondStageName:  "v",
						SecondStageIndex: 3,
					},
					{
						FirstStageName:  "1",
						SecondStageName: "w",
					},
				},
				LeftTitle: []matrix.Variable{
					{
						FirstStageName:   "y",
						FirstStageIndex:  0,
						SecondStageName:  "u",
						SecondStageIndex: 0,
					},
					{
						FirstStageName:   "y",
						FirstStageIndex:  1,
						SecondStageName:  "u",
						SecondStageIndex: 1,
					},
					{
						FirstStageName:   "y",
						FirstStageIndex:  2,
						SecondStageName:  "u",
						SecondStageIndex: 2,
					},
					{
						FirstStageName:  "z",
						SecondStageName: "1",
					},
				},
			},
			expected: DoubledOptimalSolution{
				MaxSolution: MaxSolution{
					Max: 36,
					Solution: Solution{
						Result: []float64{0, 22, 0, 8},
						Matrix: matrix.Matrix{
							Rows: []matrix.Row{
								{4, 2, 1, 0.9999999999999998, 22},
								{1.5, 0.5, 0.5, 1, 8},
								{4.5, 2.5, 1.5, 1, 25},
								{5.5, 3.5, 1.5, 1.9999999999999998, 36},
							},
							TopTitle: []matrix.Variable{
								{
									FirstStageName:   "x",
									FirstStageIndex:  0,
									SecondStageName:  "v",
									SecondStageIndex: 0,
								},
								{
									FirstStageName:   "y",
									FirstStageIndex:  0,
									SecondStageName:  "u",
									SecondStageIndex: 0,
								},
								{
									FirstStageName:   "y",
									FirstStageIndex:  2,
									SecondStageName:  "u",
									SecondStageIndex: 2,
								},
								{
									FirstStageName:   "x",
									FirstStageIndex:  2,
									SecondStageName:  "v",
									SecondStageIndex: 2,
								},
								{
									FirstStageName:  "1",
									SecondStageName: "w",
								},
							},
							LeftTitle: []matrix.Variable{
								{
									FirstStageName:   "x",
									FirstStageIndex:  1,
									SecondStageName:  "v",
									SecondStageIndex: 1,
								},
								{
									FirstStageName:   "x",
									FirstStageIndex:  3,
									SecondStageName:  "v",
									SecondStageIndex: 3,
								},
								{
									FirstStageName:   "y",
									FirstStageIndex:  1,
									SecondStageName:  "u",
									SecondStageIndex: 1,
								},
								{
									FirstStageName:  "z",
									SecondStageName: "1",
								},
							},
						},
					},
				},
				MinSolution: MinSolution{
					Min: 36,
					Solution: Solution{
						Result: []float64{3.5, 0, 1.5},
						Matrix: matrix.Matrix{
							Rows: []matrix.Row{
								{4, 2, 1, 0.9999999999999998, 22},
								{1.5, 0.5, 0.5, 1, 8},
								{4.5, 2.5, 1.5, 1, 25},
								{5.5, 3.5, 1.5, 1.9999999999999998, 36},
							},
							TopTitle: []matrix.Variable{
								{
									FirstStageName:   "x",
									FirstStageIndex:  0,
									SecondStageName:  "v",
									SecondStageIndex: 0,
								},
								{
									FirstStageName:   "y",
									FirstStageIndex:  0,
									SecondStageName:  "u",
									SecondStageIndex: 0,
								},
								{
									FirstStageName:   "y",
									FirstStageIndex:  2,
									SecondStageName:  "u",
									SecondStageIndex: 2,
								},
								{
									FirstStageName:   "x",
									FirstStageIndex:  2,
									SecondStageName:  "v",
									SecondStageIndex: 2,
								},
								{
									FirstStageName:  "1",
									SecondStageName: "w",
								},
							},
							LeftTitle: []matrix.Variable{
								{
									FirstStageName:   "x",
									FirstStageIndex:  1,
									SecondStageName:  "v",
									SecondStageIndex: 1,
								},
								{
									FirstStageName:   "x",
									FirstStageIndex:  3,
									SecondStageName:  "v",
									SecondStageIndex: 3,
								},
								{
									FirstStageName:   "y",
									FirstStageIndex:  1,
									SecondStageName:  "u",
									SecondStageIndex: 1,
								},
								{
									FirstStageName:  "z",
									SecondStageName: "1",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			support, err := FindSupportSolution(tt.m)
			if err != nil {
				t.Fatal(err)
			}

			actual, err := FindMaxDoubledWithOptimalSolution(support.Matrix)
			if err != nil {
				t.Fatal(err)
			}

			if actual == nil || !reflect.DeepEqual(*actual, tt.expected) {
				t.Fatalf("want: %v\ngot:%v", tt.expected, actual)
			}
		})
	}
}
