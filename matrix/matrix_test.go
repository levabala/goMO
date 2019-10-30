package matrix

import (
	"reflect"
	"testing"
)

func Test_shellV(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"shell1", args{length: 5}, Vector{0, 0, 0, 0, 0}},
		{"shell2", args{length: 3}, Vector{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShellV(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shellV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shellM(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{"test", args{width: 3, height: 2}, Matrix{{0, 0, 0}, {0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShellM(tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shellM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transpose(t *testing.T) {
	type args struct {
		m Matrix
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{"test1", args{Matrix{{1, 2, 3}, {4, 5, 6}}}, Matrix{{1, 4}, {2, 5}, {3, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Transpose(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_width(t *testing.T) {
	type args struct {
		m Matrix
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{Matrix{{1, 2, 3}, {4, 5, 6}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.m.Width(); got != tt.want {
				t.Errorf("width() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_height(t *testing.T) {
	type args struct {
		m Matrix
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{Matrix{{1, 2, 3}, {4, 5, 6}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.m.Height(); got != tt.want {
				t.Errorf("height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_size(t *testing.T) {
	type args struct {
		m Matrix
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"test1", args{Matrix{{1, 2, 3}, {4, 5, 6}}}, 3, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.args.m.Size()
			if got != tt.want {
				t.Errorf("size() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("size() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_sumV(t *testing.T) {
	type args struct {
		v Vector
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{Vector{1, 2, 3}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumV(tt.args.v); got != tt.want {
				t.Errorf("sumV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		m1 Matrix
		m2 Matrix
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{"test1", args{Matrix{{1, 2, 3}, {4, 5, 6}}, Matrix{{-1, -2, -3}, {-4, -5, -6}}}, Matrix{{0, 0, 0}, {0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.m1, tt.args.m2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_substract(t *testing.T) {
	type args struct {
		m1 Matrix
		m2 Matrix
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{"test1", args{Matrix{{1, 2, 3}, {4, 5, 6}}, Matrix{{1, 2, 3}, {4, 5, 6}}}, Matrix{{0, 0, 0}, {0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substract(tt.args.m1, tt.args.m2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("substract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplyWithNumber(t *testing.T) {
	type args struct {
		m     Matrix
		value float64
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{"test1", args{Matrix{{1, 2, 3}, {4, 5, 6}}, 2}, Matrix{{2, 4, 6}, {8, 10, 12}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.m.MultiplyWithNumber(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multiplyWithNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplyElementByElement(t *testing.T) {
	type args struct {
		v1 Vector
		v2 Vector
	}
	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"test1", args{Vector{1, 2, 3}, Vector{1, 2, 3}}, Vector{1, 4, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyElementByElement(tt.args.v1, tt.args.v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multiplyElementByElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiply(t *testing.T) {
	type args struct {
		m1 Matrix
		m2 Matrix
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{"test1", args{
			Matrix{
				{2, 1, 4},
				{0, 1, 1},
			},
			Matrix{
				{6, 3, -1, 0},
				{1, 1, 0, 4},
				{-2, 5, 0, 2},
			},
		}, Matrix{
			{5, 27, -2, 12},
			{-1, 6, 0, 6},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.m1, tt.args.m2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGauss(t *testing.T) {
	type args struct {
		m Matrix
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{"test1", args{
			Matrix{
				{1, -2, 1, 1},
				{2, 3, -1, -1},
				{1, -1, 2, 0},
			}},
			Matrix{
				{1.0, 0, 0, 0.20000000000000007},
				{0, 1, 0, -3.0 / 5},
				{0, 0, 1, -0.39999999999999997},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.m.Gauss(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gauss() = %v, want %v", got, tt.want)
			}
		})
	}
}
