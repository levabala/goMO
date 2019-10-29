package lpt

import (
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
	coeffs := task.targetFunction.coeffs

	if task.targetFunction.bound == BoundMax {
		coeffs = coeffs.MultiplyWithNumber(-1)
	}

	// II. Map all operators to Equal by adding new x-es (also appending new x-es to singConditions)
	limitations := make([]ConditionEqual, len(task.limitations))

	// cast non-Equal to Equal operators and add new x-es
	for i, lim := range task.limitations {
		operandsLeft := lim.operandsLeft
		operandRight := lim.operandRight
		if lim.operator != OperatorEqual {
			maxXIndex++
			newOperandsLeft := matrix.ShellV(maxXIndex)

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

			newOperandsLeft[maxXIndex-1] = x
			operandsLeft = newOperandsLeft
		}

		limitations[i] = ConditionEqual{
			operandsLeft,
			operandRight,
		}
	}

	// make every condition's operandsLeft Vector length equal
	for i, lim := range limitations {
		newOperandsLeft := matrix.ShellV(maxXIndex)
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
	limitedVector := make([]int, len(signConditions))
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
	unlimitedVector := make([]int, len(limitedVector))
	for i, v := range limitedVector {
		if v == 0 {
			unlimitedVector[i] = 1
		}
	}

	return CLPT{}
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

		limitations[i] = cond
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
