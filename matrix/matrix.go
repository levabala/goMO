package matrix

import (
	"fmt"
	"math"
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

// ShellVWithValue it's ShellV with value
func ShellVWithValue(length int, value float64) Vector {
	v := ShellV(length)

	for i := range v {
		v[i] = value
	}

	return v
}

// ShellMWithValue it's ShellN with value
func ShellMWithValue(width, height int, value float64) Matrix {
	m := make(Matrix, height)

	for y := range m {
		m[y] = make(Vector, width)
	}

	return m
}

// ShellM generates Matrix wth specified size filled with 0
func ShellM(width, height int) Matrix {
	m := make(Matrix, height)

	for y := range m {
		m[y] = make(Vector, width)
	}

	return m
}

// FindIndex finds the index of a value
func (v Vector) FindIndex(value float64) int {
	for i, val := range v {
		if val == value {
			return i
		}
	}

	return -1
}

// IsBaseVector checks if the vector is base-vector
func (v Vector) IsBaseVector() bool {
	return v.CountValue(0) == len(v)-1 && v.CountValue(1) == 1
}

// Transpose performce Matrix transposation
func (m Matrix) Transpose() Matrix {
	w, h := m.Size()
	m2 := ShellM(h, w)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			m2[x][y] = m[y][x]
		}
	}

	return m2
}

// Width returns the width of the Matrix
func (m Matrix) Width() int {
	return len(m[0])
}

// Height returns the Height of the Matrix
func (m Matrix) Height() int {
	return len(m)
}

// Size returns (width, height) of the Matrix
func (m Matrix) Size() (int, int) {
	return m.Width(), m.Height()
}

// SumV sums every Vector's value
func (v Vector) Sum() float64 {
	acc := 0.0
	for _, el := range v {
		acc += el
	}

	return acc
}

func HumaniazeValue(value float64) string {
	remain := value - math.Round(value)
	if remain < 10e-3 {
		return fmt.Sprintf("%.0f", value)
	}

	return fmt.Sprintf("%.3f", value)
}

func (v Vector) CountValue(value float64) int {
	count := 0
	for _, el := range v {
		if el == value {
			count++
		}
	}

	return count
}

// Max returns max value of the vector's value
func (v Vector) Max() float64 {
	max := 0.0
	for _, el := range v {
		max = math.Max(max, el)
	}

	return max
}

// Min returns min value of the vector's value
func (v Vector) Min() float64 {
	min := 0.0
	for _, el := range v {
		min = math.Min(min, el)
	}

	return min
}

// Max returns max value of the vector's value
func (m Matrix) Max() float64 {
	max := 0.0
	for _, row := range m {
		max = math.Max(max, row.Max())
	}

	return max
}

// Min returns min value of the vector's value
func (m Matrix) Min() float64 {
	min := 0.0
	for _, row := range m {
		min = math.Min(min, row.Min())
	}

	return min
}

// Add adds Matrixes element by element
func Add(m1, m2 Matrix) Matrix {
	w, h := m1.Size()
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
	w, h := m1.Size()
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
	mr := ShellM(m.Size())

	for y, row := range m {
		for x, element := range row {
			mr[y][x] = element * value
		}
	}

	return mr
}

// FillWith fills m1 with m2 values
func (m Matrix) FillWith(m2 Matrix) Matrix {
	m3 := m.Clone()

	for y, row := range m2 {
		for x, value := range row {
			m3[y][x] = value
		}
	}

	return m3
}

func (v Vector) FillWith(v2 Vector) Vector {
	v3 := v.Clone()

	for i, val := range v2 {
		v3[i] = val
	}

	return v3
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
func (v1 Vector) MultiplyElementByElement(v2 Vector) Vector {
	v := ShellV(len(v1))
	for i, el1 := range v1 {
		v[i] = el1 * v2[i]
	}

	return v
}

// Multiply does just Matrixes multiplying
func Multiply(m1, m2 Matrix) Matrix {
	h1 := m1.Height()
	w2 := m2.Width()

	mr := ShellM(w2, h1)
	m2Columns := m2.Transpose()

	for y, row := range m1 {
		for x := 0; x < w2; x++ {
			v := row.MultiplyElementByElement(m2Columns[x])
			mr[y][x] = v.Sum()
		}
	}

	return mr
}

// MultiplyRow multiplies each row's element by value
func MultiplyRow(m Matrix, rowIndex int, value float64) Matrix {
	w, h := m.Size()
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
func (m Matrix) DivideRow(rowIndex int, value float64) Matrix {
	w, h := m.Size()
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
func (m Matrix) SubstractRow(rowIndexWhich int, rowIndexFrom int, multiplier float64) Matrix {
	w, h := m.Size()
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
	w, h := m.Size()
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

// String converts Matrix to string
func (m Matrix) String() string {
	s := ""
	for _, row := range m {
		s += row.String()
		s += "\n"
	}
	return s
}

// String converts Vector to string
func (v Vector) String() string {
	s := ""
	for _, element := range v {
		s += fmt.Sprintf("%6.3f", element) + " "
	}
	return s
}

// BaseVector creates a base vector at provided column with 1 at provided row
func (m Matrix) BaseVector(rowIndex, columnIndex int) Matrix {
	mr := m.DivideRow(rowIndex, m[rowIndex][columnIndex])
	for y, row := range m {
		if y != rowIndex {
			mr = mr.SubstractRow(rowIndex, y, row[columnIndex])
		}
	}

	return mr
}

// Gauss makes gauss transform with the Matrix
func (m Matrix) Gauss() Matrix {
	width, height := m.Size()

	mr := m
	usedColumns := ShellV(width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if usedColumns[x] == 0 && mr[y][x] != 0 {
				mr = mr.BaseVector(y, x)
				usedColumns[x] = 1
				break
			}
		}
	}

	return mr
}

// GetColumn returns column vector at index i
func (m Matrix) GetColumn(i int) Vector {
	v := ShellV(m.Height())
	for y, row := range m {
		v[y] = row[i]
	}

	return v
}

// GetLastColumn returns last column
func (m Matrix) GetLastColumn() Vector {
	v := ShellV(m.Height())
	i := m.Width() - 1
	for y, row := range m {
		v[y] = row[i]
	}

	return v
}

// OriginalBaseVector returns original base vector
func (m Matrix) OriginalBaseVector() Matrix {
	w := m.Width()

	println("start")
	println(m.String())

	mr := m.Gauss()
	println("after gauss")
	println(mr.String())

	i := 0
	for {
		B := mr.GetLastColumn()

		everyIsPositive := true
		for _, b := range B {
			everyIsPositive = everyIsPositive && b > 0
		}

		if everyIsPositive {
			println("success")
			break
		}
		println("still not every b > 0")

		minBIndex := -1
		minBValue := 0.0
		for i, b := range B {
			if b <= 0 && minBValue > b {
				minBValue = b
				minBIndex = i
			}
		}

		fmt.Printf("working with row %d\n", minBIndex)

		for y, row := range mr {
			if row[w-1] < 0 && y != minBIndex {
				mr = mr.SubstractRow(minBIndex, y, 1)
			}
		}

		println("after substrtact")
		println(mr.String())

		mr = MultiplyRow(mr, minBIndex, -1)

		println("after mutliply -1")
		println(mr.String())

		pivotColumnIndex := -1
		for x, a := range mr[minBIndex] {
			if a > 0 && x < w-1 {
				pivotColumnIndex = x
				break
			}
		}

		pivotRowIndex := -1
		pivotValue := math.MaxFloat64
		for y, row := range mr {
			val := row[w-1] / row[pivotColumnIndex]

			if val >= 0 && val < pivotValue && (row[pivotColumnIndex] <= 0 || y == minBIndex) {
				pivotValue = val
				pivotRowIndex = y
			}
		}

		i++

		// fmt.Printf()
		// println("pre-baseVector")
		// println(mr.String())

		fmt.Printf("pivot at x: %d, y: %d\n", pivotColumnIndex, pivotRowIndex)
		mr = mr.BaseVector(pivotRowIndex, pivotColumnIndex)

		println("after base vector")
		println(mr.String())
	}

	return mr
}

// SetValue sets a value at an index
func (v Vector) SetValue(index int, value float64) Vector {
	v[index] = value
	return v
}

func (m Matrix) GetBasis() Vector {
	columns := m.Transpose()
	basis := ShellV(len(columns))

	B := m.GetLastColumn()

	for x, column := range columns {
		isBase := column.IsBaseVector()
		if isBase {
			bIndex := column.FindIndex(1)
			basis[x] = B[bIndex]
		}
	}

	return basis
}
