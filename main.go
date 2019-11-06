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

// import "gomo/matrix"

// func main() {
// 	// m := matrix.Matrix{
// 	// 	{1, -1, -1, -1, 1, 0, -46},
// 	// 	{0, -2, -1, -3, 1, 1, 4},
// 	// 	{3, 0, 2, 1, -1, -1, 2},
// 	// }

// 	m2 := matrix.Matrix{
// 		{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100},
// 		{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 200},
// 		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 300},
// 		{1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 100},
// 		{0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 100},
// 		{0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 300},
// 		{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 300},
// 	}

// 	println(m2.OriginalBaseVector().ToString())
// 	// m.TupoyPerebor()

// 	// matrix.Infinite()
// }

import (
	"gomo/lpt"
	"strings"
)

func main() {
	input := `
| 1x1 +1x2 +1x3 <= 850
| 1x4 +1x5 +1x6 <= 520
| 1x1 +1x4 = 410
| 1x2 +1x5 = 580
| 1x3 +1x6 = 350
1x1 >= 0, 1x2 >= 0, 1x3 >= 0, 1x4 >= 0, 1x5 >= 0, 1x6 >= 0
Z = 50x1 +100x2 +200x3 +160x4 +130x5 +170x6 -> (min)`
	l := lpt.ParseLPT(strings.Split(input, "\n")[1:])
	lc := lpt.CanonicalForm(l)

	m := lc.LimitationsAsMatrix().OriginalBaseVector()
	println(m.ToString())

	lcc := lc.SetMatrix(m)

	res := lcc.DoSimplex().LimitationsAsMatrix().ToString()
	println()
	println("Result:")
	println(res)
}

// func main() {
// 	input := `
// | 1x1 +1x5 +1x9 <= 100
// | 1x2 +1x6 +1x10 <= 150
// | 1x3 +1x7 +1x11 <= 120
// | 1x4 +1x8 +1x12 <= 451
// | 1x1 +1x2 +1x3 +1x4 = 219
// | 1x5 +1x6 +1x7 +1x8 = 320
// | 1x9 +1x10 +1x11 +1x12 = 230
// 1x1 >= 0, 1x2 >= 0, 1x3 >= 0, 1x4 >= 0, 1x5 >= 0, 1x6 >= 0, 1x7 >= 0, 1x8 >= 0, 1x9 >= 0, 1x10 >= 0, 1x11 >= 0, 1x12 >= 0
// Z = 19x1 +12x2 +15x3 +14x4 +15x5 +16x6 +13x7 +11x8 +21x9 +13x10 +12x11 +17x12 -> (min)`

// 	l := lpt.ParseLPT(strings.Split(input, "\n")[1:])
// 	lc := lpt.CanonicalForm(l)

// 	m := lc.LimitationsAsMatrix().OriginalBaseVector()
// 	println(m.ToString())

// 	lcc := lc.SetMatrix(m)

// 	lcc.DoSimplex()
// }
