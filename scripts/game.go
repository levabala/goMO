package scripts

import "gomo/game"
import "gomo/matrix"
import "gomo/lpt"
import "strings"

// GameBounds GameBounds
func GameBounds() {
	gameMatrix := matrix.Matrix{
		{3, 3, 2, 5},
		{4, 4, 3, 2},
		{7, 7, 4, 5},
		{4, 3, 3, 2},
		{4, 3, 4, 6},
	}

	bounds := game.GetBounds(gameMatrix)

	println(bounds.String())
}

// GameSimplify GameSimplify
func GameSimplify() {
	gameMatrix := matrix.Matrix{
		{1, 3, 4, 5},
		{4, 4, 4, 6},
		{5, 4, 3, 6},
		{4, 3, 3, 2},
		{5, 3, 4, 5},
	}

	simplified := game.Simplify(gameMatrix)
	println(gameMatrix.String())
	println("to")
	println(simplified.String())
}

// Game2x2 Game2x2
func Game2x2() {
	m := matrix.Matrix{
		{-6, 1},
		{3, -7},
	}

	solution := game.SolveGame2x2(m)
	println(solution.String())
}

// GameSolve GameSolve
func GameSolve() {
	input := `
| 13x1 +3x2 +11x3 >= 1
| 1x1 +14x2 +12x3 >= 1
| 4x1 +17x2 +2x3 >= 1
| 13x1 +14x2 +12x3 >= 1
1x1 >= 0, 1x2 >= 0, 1x3 >= 0
Z = 1x1 +1x2 +1x3 -> (min)`

	l := lpt.ParseLPT(strings.Split(input, "\n")[1:])
	ld := l.GenerateDualTask()
	ldc := ld.CanonicalForm()
	m := ldc.LimitationsAsMatrix().OriginalBaseVector()

	ldcs := ldc.SetMatrix(m).DoSimplex()

	mres := ldcs.LimitationsAsMatrix()

	println()
	println(l.String())
	println("\nto dual task")
	println(ld.String())
	println("\nto canonical form")
	println(ldc.String())
	println("\nto simplex for")
	println(ldcs.String())
	println("\nwhere matrix is")
	println(mres.String())
}
