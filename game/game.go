package game

import (
	"fmt"
	"gomo/lpt"
	"gomo/matrix"
	"math"
)

// Bound contains index and value
type Bound struct {
	index int
	value float64
}

// Bounds contains top and bottom bound
type Bounds struct {
	topBound    Bound
	bottomBound Bound
}

// Solution contains data about a game solution
type Solution struct {
	probabilities1 matrix.Vector
	probabilities2 matrix.Vector
	cost           float64
	bounds         Bounds
}

// String stringifies provided bound
func (b Bound) String() string {
	return fmt.Sprintf("index: %d, value: %f", b.index, b.value)
}

// String stringifies provided bounds
func (bs Bounds) String() string {
	return fmt.Sprintf("top:\t(%s)\nbottom:\t(%s)", bs.topBound, bs.bottomBound)
}

func (s Solution) String() string {
	return fmt.Sprintf("ps1:\t[%s]\nps2:\t[%s]\ncost:\t%f\nbounds:\n%s", s.probabilities1, s.probabilities2, s.cost, s.bounds)
}

// GetBounds calcs alpha and betta bounds
func GetBounds(m matrix.Matrix) Bounds {
	// w := m.Width() + 1
	// h := m.Height() + 1
	// mAug := matrix.ShellM(w, h).FillWith(m)

	// for _, row := range mAug {
	// 	row[w - 1] = row.Max()
	// }

	// for x := 0; x < w; x++ {
	// 	column := m.GetColumn(x)
	// 	mAug[h - 1][x] = column.Max()
	// }

	w := m.Width()

	alphaIndex := -1
	alphaValue := -math.MaxFloat64
	for y, row := range m {
		v := row.Max()
		if v > alphaValue {
			alphaIndex = y
			alphaValue = v
		}
	}

	betaIndex := -1
	betaValue := math.MaxFloat64
	for x := 0; x < w; x++ {
		column := m.GetColumn(x)
		v := column.Max()

		if v < betaValue {
			betaIndex = x
			betaValue = v

			if betaIndex == alphaIndex {
				break
			}
		}
	}

	return Bounds{
		Bound{
			alphaIndex,
			alphaValue,
		},
		Bound{
			betaIndex,
			betaValue,
		},
	}
}

// Simplify simplifies given matrix
func Simplify(m matrix.Matrix) matrix.Matrix {
	w, h := m.Size()
	removedColumns := matrix.ShellV(w)
	removedRows := matrix.ShellV(h)

	for y1, row1 := range m {
		if removedRows[y1] == 1 {
			continue
		}

		for y2, row2 := range m {
			if removedRows[y2] == 1 || y2 == y1 {
				continue
			}

			everyIsLess := true
			for x, el1 := range row1 {
				everyIsLess = everyIsLess && (removedColumns[x] == 1 || el1 >= row2[x])
			}

			if everyIsLess {
				removedRows[y2] = 1
			}
		}
	}

	columns := m.Transpose()

	for x1, column1 := range columns {
		if removedColumns[x1] == 1 {
			continue
		}

		for x2, column2 := range columns {
			if removedColumns[x2] == 1 || x2 == x1 {
				continue
			}

			everyIsBigger := true
			for y, el1 := range column1 {
				everyIsBigger = everyIsBigger && (removedRows[y] == 1 || el1 <= column2[y])
			}

			if everyIsBigger {
				removedRows[x2] = 1
			}
		}
	}

	removedColumnsCount := removedColumns.CountValue(1)
	removedRowsCount := removedRows.CountValue(1)

	matrixNew := matrix.ShellM(w-removedColumnsCount, h-removedRowsCount)

	realY := 0
	for y, row := range m {
		if removedRows[y] == 1 {
			continue
		}

		realX := 0
		for x, el := range row {
			if removedColumns[x] == 1 {
				continue
			}

			matrixNew[realY][realX] = el
			realX++
		}

		realY++
	}

	return matrixNew
}

// SolveGame2x2 solves game2x2
func SolveGame2x2(m matrix.Matrix) Solution {
	a11 := m[0][0]
	a21 := m[1][0]
	a12 := m[0][1]
	a22 := m[1][1]

	p1 := (a22 - a12) / (a11 - a21 - a12 + a22)
	p2 := 1 - p1

	q1 := (a22 - a21) / (a11 - a12 - a21 + a22)
	q2 := 1 - q1

	v := p1*a11 + p2*a12

	bounds := GetBounds(m)

	solution := Solution{
		probabilities1: matrix.Vector{p1, p2},
		probabilities2: matrix.Vector{q1, q2},
		cost:           v,
		bounds:         bounds,
	}

	return solution
}

// SolveGame solves game mxn
func SolveGame(m matrix.Matrix) Solution {
	m = m.Clone().Transpose()
	minValue := m.Min()
	wOriginal, hOriginal := m.Size()

	appendix := 0.0
	if minValue < 0 {
		appendix = minValue*-1 + 1

		for y, row := range m {
			for x, el := range row {
				m[y][x] = el + appendix
			}
		}
	}

	lptMatrix := matrix.ShellM(wOriginal+1, hOriginal).FillWith(m)
	w, _ := lptMatrix.Size()

	for _, row := range lptMatrix {
		row[w-1] = 1
	}

	println("Start matrix:")
	println(lptMatrix.String())
	println()

	println("With appendix:")
	println(lptMatrix.String())
	println()

	operators := make([]lpt.Operator, len(lptMatrix))
	for i := range operators {
		operators[i] = lpt.OperatorLessOrEqual
	}

	l := lpt.LPT{}.
		SetMatrix(lptMatrix, operators).
		SetSignConditionToEvery(lpt.OperatorGreaterOrEqual).
		SetDefaultTargetFunction()

	println("Resulting LPT:")
	println(l.String())
	println()

	ld := l.GenerateDualTask()

	println("Dual LPT:")
	println(ld.String())
	println()

	ldc := ld.CanonicalForm()

	println("Canonical LPT:")
	println(ldc.String())
	println()

	ml := ldc.LimitationsAsMatrix().OriginalBaseVector()
	ldcs, zValues := ldc.SetMatrix(ml).DoSimplex()
	mlres := ldcs.LimitationsAsMatrix()

	println()
	println(mlres.String())

	basis2 := mlres.GetBasis()
	println(basis2.String())

	valuesForY := basis2[:len(zValues)-hOriginal]
	basis2WithZeros := make(matrix.Vector, len(basis2)).FillWith(valuesForY)

	gameCost := 1 / ldcs.ToLPT().MutliplyTargetFunctionWith(basis2WithZeros).Sum()

	probabilities2 := valuesForY.MultiplyWithNumber(gameCost)

	probabilities1 := zValues[len(zValues)-hOriginal : len(zValues)-1].MultiplyWithNumber(gameCost)

	return Solution{
		probabilities1: probabilities1,
		probabilities2: probabilities2,
		cost:           gameCost - appendix,
	}
}
