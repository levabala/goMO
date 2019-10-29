package main

import (
	"gomo/lpt"
	"strings"
)

func main() {
	input := `
| 1x1 -1x2 >= -2
| 5x1 +2x2 <= 15
| 3x1 -1x2 -1x3 = 3
1x2 >= 0, 1x3 >= 0
Z = 1x1 -2x3 -> (max)`

	l := lpt.ParseLPT(strings.Split(input, "\n")[1:])
	lpt.CanonicalForm(l)
}
