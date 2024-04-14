package games

import (
	"math/rand"

	"github.com/hrvadl/algo/internal/matrix"
)

type (
	Weight  = float64
	Weights = []Weight
)

type SimulationStep struct {
	FirstPlayerRandomNum float64
	FirstPlayerStrategy  matrix.Variable

	SecondPlayerRandomNum float64
	SecondPlayerStrategy  matrix.Variable

	FirstPlayerWin    float64
	FirstPlayerWinSum float64
	FirstPlayerWinAvg float64
}

type SimulationOptions struct {
	Matrix               matrix.Matrix
	FirstPlayerStrategy  Weights
	SecondPlayerStrategy Weights
	Times                int
}

func SimulateGame(opt SimulationOptions) []SimulationStep {
	total := make([]SimulationStep, 0, opt.Times)
	var totalWin float64

	for i := range opt.Times {
		fn := rand.Float64()
		sn := rand.Float64()
		fs := GetStrategyFromNum(opt.FirstPlayerStrategy, fn)
		ss := GetStrategyFromNum(opt.SecondPlayerStrategy, sn)
		win := opt.Matrix.Rows[fs][ss]
		totalWin += win

		total = append(total, SimulationStep{
			FirstPlayerRandomNum: fn,
			FirstPlayerStrategy: matrix.Variable{
				FirstStageName:  "x",
				FirstStageIndex: fs,
			},
			SecondPlayerRandomNum: sn,
			SecondPlayerStrategy: matrix.Variable{
				SecondStageName:  "y",
				SecondStageIndex: ss,
			},
			FirstPlayerWin:    win,
			FirstPlayerWinSum: totalWin,
			FirstPlayerWinAvg: totalWin / float64((i + 1)),
		})
	}

	return total
}

func GetStrategyFromNum(weights Weights, n float64) int {
	var sumPrev float64
	for i, w := range weights {
		if w == 0 {
			continue
		}

		if n <= w+sumPrev {
			return i
		}

		sumPrev += w
	}

	return 0
}
