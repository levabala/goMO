package scripts

import "gomo/game"
import "gomo/matrix"

// Bounds Bounds
func Bounds() {
	gameMatrix := matrix.Matrix{
		{3, 3, 2, 5},
		{4, 4, 3, 2},
		{7, 7, 4, 5},
		{4, 3, 3, 2},
		{4, 3, 4, 6},
	}

	bounds := game.GetBounds(gameMatrix)
	index := bounds.TopBound.index

	println(index)
}

// Simplify Simplify
func Simplify() {
	gameMatrix := matrix.Matrix{
		{1, 3, 4, 5},
		{4, 4, 4, 6},
		{5, 4, 3, 6},
		{4, 3, 3, 2},
		{5, 3, 4, 5},
	}

	simplified := game.Simplify(gameMatrix)
	println(gameMatrix.ToString())
	println("to")
	println(simplified.ToString())
}
