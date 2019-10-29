package lpt

import "gomo/matrix"


type Bound int
const (
	BoundMin Bound = iota
	BoundMax Bound = iota
)

type Operator int
const (
	OperatorGreater Operator = iota
	OperatorGreaterOrEqual Operator = iota
	OperatorLess Operator = iota
	OperatorLessOrEqual Operator = iota
	OperatorEqual Operator = iota
)

type TargetFunction struct {
	coeffs matrix.Vector
	bound Bound
}

type Condition struct {
	operandsLeft matrix.Vector
	operator Operator
	operandRight float64
}

type ConditionZero struct {
	operandsLeft matrix.Vector
	operator Operator
}

type LPT struct {
	limitations []Condition
	singConditions []ConditionZero
	targetFunction TargetFunction
}


// specific types for LPTC (Lineral Programming Tasks Canonical)
type TargetFunctionMin struct {
	coeffs matrix.Vector
}

type ConditionEqual struct {
	operandsLeft matrix.Vector
	operandRight float64
}

type ConditionZeroPositive struct {
	operandsLeft matrix.Vector
}

type LPTC struct {
	limitations []ConditionEqual
	singConditions []ConditionZeroPositive
	targetFunction TargetFunctionMin
}

func CanonicalForm(task LPT) LPTC {
	// I. Minimize target function
	coeffs := task.targetFunction.coeffs

	if task.targetFunction.bound == BoundMax {
		coeffs = coeffs.MultiplyWithNumber(-1)
	}

	// II. Map all operators to Equal by adding new x-es (also appending new x-es to singConditions)
	limitations := make([]ConditionEqual, len(task.limitations))

	newXesCount := 0
	for i, lim := range task.limitations {		
		operandsLeft := lim.operandsLeft
		operandRight := lim.operandRight
		if lim.operator != OperatorEqual {
			var x float64
			if lim.operator == OperatorLessOrEqual {
				x = 1
			} else {
				x = -1
			}

			operandsLeft = append(operandsLeft, x)
			newXesCount++
		}

		limitations[i] = ConditionEqual {
			operandsLeft,
			operandRight,
		}
	}

	signConditions := make([])
}