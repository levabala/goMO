package matrix

import (
	"math"
	"strconv"
)

// Vector is just 1d array of float64
type Vector []float64

// Matrix is just array of Vectors
type Matrix []Vector

// ShellV generates Vector filled with 0 with specified length
func ShellV(length int) Vector {
	v := make(Vector, length)
	return v
}

// ShellM generates Matrix wth specified size filled with 0
func ShellM(width, height int) Matrix {
	m := make(Matrix, height)

	for y := range m {
		m[y] = make(Vector, width)
	}

	return m
}

// Transpose performce Matrix transposation
func Transpose(m Matrix) Matrix {
	w, h := Size(m)
	m2 := ShellM(h, w)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			m2[x][y] = m[y][x]
		}
	}

	return m2
}

// Width returns the width of the Matrix
func Width(m Matrix) int {
	return len(m[0])
}

// Height returns the Height of the Matrix
func Height(m Matrix) int {
	return len(m)
}

// Size returns (width, height) of the Matrix
func Size(m Matrix) (int, int) {
	return Width(m), Height(m)
}

// SumV sums every Vector's value
func SumV(v Vector) float64 {
	acc := 0.0
	for _, el := range v {
		acc += el
	}

	return acc
}

// Add adds Matrixes element by element
func Add(m1, m2 Matrix) Matrix {
	w, h := Size(m1)
	m := ShellM(w, h)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			m[y][x] = m1[y][x] + m2[y][x]
		}
	}

	return m
}

// Substract substracts Matrixes element by element
func Substract(m1, m2 Matrix) Matrix {
	w, h := Size(m1)
	m := make(Matrix, h)

	for y := range m {
		m[y] = make(Vector, w)

		for x := 0; x < h; x++ {
			m[y][x] = m1[y][x] - m2[y][x]
		}
	}

	return m
}

// MultiplyWithNumber multiplies each Matrix's element by value
func (m Matrix) MultiplyWithNumber(value float64) Matrix {
	mr := ShellM(Size(m))

	for y, row := range m {
		for x, element := range row {
			mr[y][x] = element * value
		}
	}

	return mr
}

// MultiplyWithNumber multiplies each Vectors's element by value
func (v Vector) MultiplyWithNumber(value float64) Vector {
	l := len(v)
	vr := ShellV(l)

	for i := 0; i < l; i++ {
		vr[i] = v[i] * value
	}

	return vr
}

// MultiplyElementByElement multiplies two vectors element by element
func MultiplyElementByElement(v1, v2 Vector) Vector {
	v := ShellV(len(v1))
	for i, el1 := range v1 {
		v[i] = el1 * v2[i]
	}

	return v
}

// Multiply does just Matrixes multiplying
func Multiply(m1, m2 Matrix) Matrix {
	h1 := Height(m1)
	w2 := Width(m2)

	mr := ShellM(w2, h1)
	m2Columns := Transpose(m2)

	for y, row := range m1 {
		for x := 0; x < w2; x++ {
			v := MultiplyElementByElement(row, m2Columns[x])
			mr[y][x] = SumV(v)
		}
	}

	return mr
}

// MultiplyRow multiplies each row's element by value
func MultiplyRow(m Matrix, rowIndex int, value float64) Matrix {
	w, h := Size(m)
	mr := ShellM(w, h)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if y == rowIndex {
				mr[y][x] = m[y][x] * value
			} else {
				mr[y][x] = m[y][x]
			}
		}
	}

	return mr
}

// DivideRow divides each row's element by value
func DivideRow(m Matrix, rowIndex int, value float64) Matrix {
	w, h := Size(m)
	mr := ShellM(w, h)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if y == rowIndex {
				mr[y][x] = m[y][x] / value
			} else {
				mr[y][x] = m[y][x]
			}
		}
	}

	return mr
}

// SubstractRow substracts rowIndexWhich values mutliplied by multiplier from rowIndexFrom
func SubstractRow(m Matrix, rowIndexWhich int, rowIndexFrom int, multiplier float64) Matrix {
	w, h := Size(m)
	mr := ShellM(w, h)

	rowMultiplied := m[rowIndexWhich].Clone().MultiplyWithNumber(multiplier)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if y == rowIndexFrom {
				mr[y][x] = m[y][x] - rowMultiplied[x]
			} else {
				mr[y][x] = m[y][x]
			}
		}
	}

	return mr
}

// Clone clones the Matrix
func (m Matrix) Clone() Matrix {
	w, h := Size(m)
	mr := ShellM(w, h)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			mr[y][x] = m[y][x]
		}
	}

	return mr
}

// Clone clones the Vector
func (v Vector) Clone() Vector {
	l := len(v)
	vr := ShellV(l)

	for i := 0; i < l; i++ {
		vr[i] = v[i]
	}

	return vr
}

func minInt(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}

	return i2
}

// ToString converts Matrix to string
func (m Matrix) ToString() string {
	s := ""
	for _, row := range m {
		for _, element := range row {
			s += strconv.FormatFloat(element, 'f', 1, 64) + " "
		}
		s += "\n"
	}
	return s
}

// BaseVector creates a base vector at provided column with 1 at provided row
func (m Matrix) BaseVector(rowIndex, columnIndex int) Matrix {
	mr := DivideRow(m, rowIndex, m[rowIndex][columnIndex])
	for y, row := range m {
		if y != rowIndex {
			mr = SubstractRow(mr, rowIndex, y, row[columnIndex])
		}
	}

	return mr
}

// Gauss makes gauss transform with the Matrix
func (m Matrix) Gauss() Matrix {
	width, height := Size(m)
	minSide := minInt(width, height)

	mr := m
	for i := 0; i < minSide; i++ {
		mr = mr.BaseVector(i, i)
	}

	return mr
}

// GetColumn returns column vector at index i
func (m Matrix) GetColumn(i int) Vector {
	v := ShellV(Height(m))
	for y, row := range m {
		v[y] = row[i]
	}

	return v
}

// GetLastColumn returns last column
func (m Matrix) GetLastColumn() Vector {
	v := ShellV(Height(m))
	i := Width(m) - 1
	for y, row := range m {
		v[y] = row[i]
	}

	return v
}

// func (m Matrix) AllBasicSolutions() Matrix {

// }

// OriginalBaseVector returns original base vector
func (m Matrix) OriginalBaseVector() Matrix {
	w := Width(m)

	mr := m.Gauss()
	B := mr.GetLastColumn()

	minBIndex := -1
	minBValue := 0.0
	for i, b := range B {
		if b <= 0 && minBValue > b {
			minBValue = b
			minBIndex = i
		}
	}

	for y, row := range mr {
		if row[w-1] < 0 {
			mr = SubstractRow(mr, minBIndex, y, 1)
		}
	}

	mr = MultiplyRow(mr, minBIndex, -1)

	pivotColumnIndex := -1
	for x, a := range mr[minBIndex] {
		if a > 0 {
			pivotColumnIndex = x
			break
		}
	}

	pivotRowIndex := -1
	pivotValue := math.MaxFloat64
	for y, row := range mr {
		val := row[w-1] / row[pivotColumnIndex]

		if val < pivotValue {
			pivotValue = val
			pivotRowIndex = y
		}
	}

	mr = mr.BaseVector(pivotRowIndex, pivotColumnIndex)

	println(mr.ToString())

	return mr
}
