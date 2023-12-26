// # Reverse Polish Expression
//
// This package is a poorly implementation of a polish expression evaluator.

package pexpr

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/jgardona/cmath"
)

const (
	sum            = "+"
	subtraction    = "-"
	multiplication = "*"
	division       = "/"
	sine           = "sin"
	cosine         = "cos"
	logarithm      = "ln"
	exponential    = "exp"
	squareroot     = "sqrt"
)

var (
	ErrWrongSolution  = errors.New("the solution must have only one answere")
	ErrBadFunction    = errors.New("function not implemented")
	ErrDivisionByZero = errors.New("cant divide by zero")
)

// PolishEvaluator is a reverse notation polish expression evaluator.
type PolishEvaluator struct {
	expression string
	variables  []float64
}

// NewPolishEvaluator instatiates a NewPolishEvaluator.
//
// expr is the expression to be evaluated, and vars is a float array
// of variables, which is used in expression following the index order.
// Arguments of expression can be writen as $var_number. Ex. $0.
func NewPolishEvaluator(expr string, vars []float64) PolishEvaluator {
	return PolishEvaluator{expression: expr, variables: vars}
}

// Evaluate evaluates an expression.
// The returned result is a float, or an error if fail.
func (pe PolishEvaluator) Evaluate() (float64, error) {
	tokens := strings.Split(pe.expression, " ")
	arguments := make([]float64, 0, 3)

	for _, token := range tokens {
		if unicode.IsDigit(rune(token[0])) {
			if result, err := strconv.ParseFloat(token, 64); err != nil {
				return 0.0, err
			} else {
				cmath.Push(&arguments, result)
			}
		} else if strnum, ok := strings.CutPrefix(token, "$"); ok {
			if index, err := strconv.ParseInt(strnum, 10, 64); err != nil {
				return 0.0, err
			} else {
				cmath.Push(&arguments, pe.variables[index])
			}
		} else {

			value := cmath.Pop(&arguments)
			switch token {
			case sum:
				cmath.Push(&arguments, cmath.Pop(&arguments)+value)
			case subtraction:
				cmath.Push(&arguments, cmath.Pop(&arguments)-value)
			case multiplication:
				cmath.Push(&arguments, cmath.Pop(&arguments)*value)
			case division:
				if value == 0 {
					return 0.0, ErrDivisionByZero
				}
				cmath.Push(&arguments, cmath.Pop(&arguments)/value)
			case sine:
				cmath.Push(&arguments, math.Sin(value))
			case cosine:
				cmath.Push(&arguments, math.Cos(value))
			case logarithm:
				cmath.Push(&arguments, math.Log(value))
			case exponential:
				cmath.Push(&arguments, math.Exp(value))
			case squareroot:
				cmath.Push(&arguments, math.Sqrt(value))

			default:
				return 0.0, ErrBadFunction
			}
		}
	}

	if len(arguments) != 1 {
		return 0.0, ErrWrongSolution
	} else {
		return arguments[0], nil
	}
}
