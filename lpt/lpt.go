package lpt

import "gomo/matrix"

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
		l := len(lim.operandsLeft)
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
		}

		limitations[i] = ConditionEqual{
			operandsLeft,
			operandRight,
		}
	}

	// make every condition's operandsLeft Vector length equal
	for i, lim := range task.limitations {
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

		operandsLeft := matrix.ShellV(maxXIndex)
		operandsLeft[xIndex] = 1

		signConditions[i+1] = ConditionZeroPositive{
			operandsLeft,
		}
	}

	return CLPT{}
}
