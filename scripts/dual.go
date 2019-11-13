package scripts

import (
	"gomo/lpt"
	"strings"
)

// DualScript DualScript
func DualScript() {
	input := `
| 2x1 -4x4 >= 5
| 3x2 -2x3 = 10
| 4x1 +5x2 -3x4 >= 7
1x1 >= 0, 1x2 >= 0, 1x3 >= 0
Z = 3x1 +2x2 -3x3 +5x4 -> (min)`
	l := lpt.ParseLPT(strings.Split(input, "\n")[1:])

	println(input[1:])
	println("\nto\n")
	println(l.GenerateDualTask().ToString())
}
