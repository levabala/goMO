package scripts

import (
	"gomo/lpt"
	"strings"
)

// SimplexScript SimplexScript
func SimplexScript() {
	input := `
| 1x1 +1x2 +1x3 <= 850
| 1x4 +1x5 +1x6 <= 520
| 1x1 +1x4 = 410
| 1x2 +1x5 = 580
| 1x3 +1x6 = 350
1x1 >= 0, 1x2 >= 0, 1x3 >= 0, 1x4 >= 0, 1x5 >= 0, 1x6 >= 0
Z = 50x1 +100x2 +200x3 +160x4 +130x5 +170x6 -> (min)`
	l := lpt.ParseLPT(strings.Split(input, "\n")[1:])
	lc := l.CanonicalForm()

	m := lc.LimitationsAsMatrix().OriginalBaseVector()
	println(m.String())

	lcc := lc.SetMatrix(m)

	lccs, _ := lcc.DoSimplex()
	res := lccs.LimitationsAsMatrix().String()
	println()
	println("Result:")
	println(res)
}
