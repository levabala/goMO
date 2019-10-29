package main

// import (
// 	"gomo/lpt"
// 	"strings"
// )

// func main() {
// 	input := `
// | 1x1 -1x2 >= -2
// | 5x1 +2x2 <= 15
// | 3x1 -1x2 -1x3 = 3
// 1x2 >= 0, 1x3 >= 0
// Z = 1x1 -2x3 -> (max)`

// 	l := lpt.ParseLPT(strings.Split(input, "\n")[1:])
// 	lpt.CanonicalForm(l)
// }

import "gomo/matrix"

func main() {
	// m := matrix.Matrix{
	// 	{1, -1, -1, -1, 1, 0, -46},
	// 	{0, -2, -1, -3, 1, 1, 4},
	// 	{3, 0, 2, 1, -1, -1, 2},
	// }

	m2 := matrix.Matrix{
		{1, -1, -2, 0, -1, -1, 0, 2},
		{0, 1, -3, -1, -1, 0, -1, -26},
		{1, 0, 1, 1, 0, -1, -1, 17},
	}

	m2.OriginalBaseVector()
	// m.TupoyPerebor()
}
