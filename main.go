package main

import "gomo/matrix"

func main() {
	m := matrix.Matrix{
		{1, -2, 1, 1},
		{2, 3, -1, -1},
		{1, -1, 2, 0},
	}

	mg := m.OriginalBaseVector()
	println(mg.ToString())
}
