package lpt

import (
	"fmt"
	"gomo/matrix"
	"strconv"
	"strings"
)

// Bound shows or Max value or Min value
type Bound int

const (
	// BoundMin is minimal bound
	BoundMin Bound = iota
	// BoundMax is maximal bound
	BoundMax Bound = iota
)

// Operator shows different conditional operators
type Operator int

const (
	// OperatorGreater is >
	OperatorGreater Operator = iota
	// OperatorGreaterOrEqual is >=
	OperatorGreaterOrEqual Operator = iota
	// OperatorLess is <
	OperatorLess Operator = iota
	// OperatorLessOrEqual is <=
	OperatorLessOrEqual Operator = iota
	// OperatorEqual is =
	OperatorEqual Operator = iota
)

// TargetFunction is Z
type TargetFunction struct {
	coeffs matrix.Vector
	bound  Bound
}

// Condition is structure that implements such math notaion: 1x_1 + x_2 - 3x_3 >= 18
type Condition struct {
	operandsLeft matrix.Vector
	operator     Operator
	operandRight float64
}

// LPT is container for Lineral Programming Task
type LPT struct {
	limitations    []Condition
	signConditions []ConditionZero
	targetFunction TargetFunction
}

// the following are specific types for LPTC (Lineral Programming Tasks Canonical)

// TargetFunctionMin is TargetFunction where bound is locked to BoundMin
type TargetFunctionMin struct {
	coeffs matrix.Vector
}

// ConditionEqual is Conditional where operator is locked to Equal
type ConditionEqual struct {
	operandsLeft matrix.Vector
	operandRight float64
}

// ConditionZeroPositive is Conditional where operator is locked to Greater and operandRight to 0
type ConditionZeroPositive struct {
	operandsLeft matrix.Vector
}

// ConditionZero is Condition where right part is locked to 0
type ConditionZero struct {
	operandsLeft matrix.Vector
	operator     Operator
}

// CLPT is container for Canonical Lineral Programming Task
type CLPT struct {
	limitations    []ConditionEqual
	signConditions []ConditionZeroPositive
	targetFunction TargetFunctionMin
}

// CanonicalForm transforms LPT to CLPT
func CanonicalForm(task LPT) CLPT {
	// variable that shows maximal x's index (starting from 0)
	maxXIndex := 0
	for _, lim := range task.limitations {
		l := len(lim.operandsLeft) - 1
		if l > maxXIndex {
			maxXIndex = l
		}
	}

	maxXIndexAtStart := maxXIndex

	// I. Minimize target function
	targetFunctionCoeffs := task.targetFunction.coeffs

	if task.targetFunction.bound == BoundMax {
		targetFunctionCoeffs = targetFunctionCoeffs.MultiplyWithNumber(-1)
	}

	// II. Map all operators to Equal by adding new x-es (also appending new x-es to singConditions)
	limitations := make([]ConditionEqual, len(task.limitations))

	// cast non-Equal to Equal operators and add new x-es
	for i, lim := range task.limitations {
		operandsLeft := lim.operandsLeft
		operandRight := lim.operandRight
		if lim.operator != OperatorEqual {
			maxXIndex++
			newOperandsLeft := matrix.ShellV(maxXIndex + 1)

			// fill with already existing x-es
			for i, x := range operandsLeft {
				newOperandsLeft[i] = x
			}

			var x float64
			if lim.operator == OperatorLessOrEqual {
				x = 1
			} else {
				x = -1
			}

			newOperandsLeft[maxXIndex] = x
			operandsLeft = newOperandsLeft
		}

		limitations[i] = ConditionEqual{
			operandsLeft,
			operandRight,
		}
	}

	// make every condition's operandsLeft Vector length equal
	for i, lim := range limitations {
		newOperandsLeft := matrix.ShellV(maxXIndex + 1)
		for i, x := range lim.operandsLeft {
			newOperandsLeft[i] = x
		}

		limitations[i].operandsLeft = newOperandsLeft
	}

	newXesCount := maxXIndex - maxXIndexAtStart
	signConditions := make([]ConditionZeroPositive, len(task.signConditions)+newXesCount)

	// TODO: transform non-ZeroPositive conditions to ZeroPositive
	// copying already existing signConditions to new variable
	for i, el := range task.signConditions {
		signConditions[i] = ConditionZeroPositive{
			operandsLeft: el.operandsLeft,
		}
	}

	// adding new sign conditions
	for i := 0; i < newXesCount; i++ {
		xIndex := maxXIndexAtStart + i + 1

		operandsLeft := matrix.ShellV(maxXIndex + 1)
		operandsLeft[xIndex] = 1

		pushI := len(task.signConditions) + i
		signConditions[pushI] = ConditionZeroPositive{
			operandsLeft,
		}
	}

	// III. Emulate positiviness condition for unlimited variables
	limitedVector := make([]int, maxXIndex+1)
	for _, cond := range signConditions {
		xIndex := -1
		for i, v := range cond.operandsLeft {
			if v == 1 {
				xIndex = i
				break
			}
		}

		limitedVector[xIndex] = 1
	}

	// erm, it's so many codelines in golang to just invert vector's values! disappointing..
	unlimitedVector := make([]bool, len(limitedVector))
	for i, v := range limitedVector {
		unlimitedVector[i] = v == 0
	}

	// no we need to set to 0 each unlimited X and then add X' and X'' with same pre-coeff
	for i, isUnLimited := range unlimitedVector {
		if isUnLimited {
			// remove coeffs from targetFunction
			targetFunctionCoeffs[i] = 0

			// append X' and X'' to targetFunction
			targetFunctionCoeffs = append(targetFunctionCoeffs, 1, -1)

			// X' condition
			condX1V := matrix.ShellV(maxXIndex + 2)
			condX1V[maxXIndex+1] = 1

			condX1 := ConditionZeroPositive{
				condX1V,
			}

			// X'' condition
			condX2V := matrix.ShellV(maxXIndex + 3)
			condX2V[maxXIndex+2] = -1

			condX2 := ConditionZeroPositive{
				condX2V,
			}

			// append X' and X'' to sign condtions list
			signConditions = append(signConditions, condX1, condX2)

			// process limitations
			for i2, lim := range limitations {
				// take X pre-coeff
				coeff := lim.operandsLeft[i]

				// remove X
				lim.operandsLeft[i] = 0

				// append X' and X''
				lim.operandsLeft = append(lim.operandsLeft, coeff, coeff)

				// write to limitations
				limitations[i2] = lim
			}
		}
	}

	targetFunction := TargetFunctionMin{
		targetFunctionCoeffs,
	}

	return CLPT{
		limitations:    limitations,
		signConditions: signConditions,
		targetFunction: targetFunction,
	}
}

func parseX(str string) (int, int) {
	arr := strings.Split(str, "x")
	value, _ := strconv.ParseInt(arr[0], 10, 64)
	index, _ := strconv.ParseInt(arr[1], 10, 64)

	valueI := int(value)
	indexI := int(index) - 1

	return valueI, indexI
}

func parseXes(str []string) matrix.Vector {
	v := matrix.Vector{}

	for _, s := range str {
		value, index := parseX(s)

		for len(v) < index+1 {
			v = append(v, 0)
		}

		v[index] = float64(value)
	}

	return v
}

// ParseLPT parses string array to LPT
func ParseLPT(lines []string) LPT {
	linesCount := len(lines)

	limitationsS := lines[:linesCount-2]
	signConditionsS := lines[linesCount-2]
	targetFunctionS := lines[linesCount-1]

	maxXIndex := -1

	// parsing limitations
	limitations := make([]Condition, len(limitationsS))
	for i, line := range limitationsS {
		chunks := strings.Split(line, " ")
		chunksCount := len(chunks)

		operandRight, _ := strconv.ParseFloat(chunks[chunksCount-1], 64)
		operatorS := chunks[chunksCount-2]

		var operator Operator
		switch operatorS {
		case ">=":
			operator = OperatorGreaterOrEqual
		case "=":
			operator = OperatorEqual
		case "<=":
			operator = OperatorLessOrEqual
		case ">":
			operator = OperatorGreater
		case "<":
			operator = OperatorLess
		}

		coeffsS := chunks[1 : chunksCount-2]
		operandsLeft := matrix.Vector{}

		for _, coeffS := range coeffsS {
			value, index := parseX(coeffS)

			for len(operandsLeft) < index+1 {
				operandsLeft = append(operandsLeft, 0)
			}

			operandsLeft[index] = float64(value)
		}

		cond := Condition{
			operandsLeft,
			operator,
			operandRight,
		}

		if len(operandsLeft) > maxXIndex {
			maxXIndex = len(operandsLeft) - 1
		}

		limitations[i] = cond
	}

	// make every condition's operandsLeft Vector length equal
	for i, lim := range limitations {
		newOperandsLeft := matrix.ShellV(maxXIndex + 1)
		for i, x := range lim.operandsLeft {
			newOperandsLeft[i] = x
		}

		limitations[i].operandsLeft = newOperandsLeft
	}

	// parsing signs
	signConditionsSChunks := strings.Split(signConditionsS, ", ")
	signConditions := make([]ConditionZero, len(signConditionsSChunks))

	for i, chunk := range signConditionsSChunks {
		arr := strings.Split(chunk, " >= ")
		left := arr[0]

		_, xIndex := parseX(left)

		operandsLeft := matrix.ShellV(int(xIndex) + 1)
		operandsLeft[xIndex] = 1

		cond := ConditionZero{
			operandsLeft,
			OperatorGreaterOrEqual,
		}

		signConditions[i] = cond
	}

	// parsing target function
	targetFunctionSChunks1 := strings.Split(targetFunctionS, " -> ")
	targetFunctionSChunks2 := strings.Split(targetFunctionSChunks1[0], " = ")

	targetFunctionCoeffsS := strings.Split(targetFunctionSChunks2[1], " ")
	coeffs := parseXes(targetFunctionCoeffsS)

	var bound Bound
	if targetFunctionSChunks1[1] == "(max)" {
		bound = BoundMax
	} else {
		bound = BoundMin
	}

	targetFunction := TargetFunction{
		coeffs,
		bound,
	}

	l := LPT{
		limitations:    limitations,
		signConditions: signConditions,
		targetFunction: targetFunction,
	}

	return l
}

// LimitationsAsMatrix returns tasks' limitations in Matrix form
func (task CLPT) LimitationsAsMatrix() matrix.Matrix {
	m := matrix.ShellM(len(task.limitations[0].operandsLeft)+1, len(task.limitations))

	for i, lim := range task.limitations {
		for x, value := range lim.operandsLeft {
			// place x-es
			m[i][x] = value
		}

		// place b also
		m[i][len(lim.operandsLeft)] = lim.operandRight
	}

	return m
}

// func (task CLPT) SetMatrix(m matrix.Matrix) CLPT {
// 	limitations := make([]ConditionEqual, len(task.limitations))
// }

// DoSimplex performs Simplex transformation
func DoSimplex(task CLPT) CLPT {
	m := task.LimitationsAsMatrix()
	_, h := m.Size()

	baseVector := make([]float64, h)

	columns := matrix.Transpose(m)
	for x, column := range columns {
		zerosCount := 0
		onesCount := 0
		onePosition := -1
		for y, el := range column {
			if el == 0.0 {
				zerosCount++
			} else if el == 1.0 {
				onesCount++
				onePosition = y
			}
		}

		isBase := onesCount == 1 && zerosCount == h-1
		if isBase {
			baseVector[onePosition] = task.targetFunction.coeffs[x]
		}
	}

	B := m.GetLastColumn()
	totalZ := matrix.SumV(matrix.MultiplyElementByElement(B, baseVector))

	fmt.Printf("%f\n", totalZ)

	return task
}
