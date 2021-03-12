// Package calculator provides a library for
// simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the
// result of adding them together
func Add(inputs ...float64) float64 {
	var result float64 = 0
	for _, input := range inputs {
		result += input
	}
	return result
}

// Substract takes two numbers and returns the
// result of substracting the first from the last
func Substract(inputs ...float64) float64 {
	var result float64 = 0
	for i, input := range inputs {
		if i == 0 {
			result = input
		} else {
			result -= input
		}
	}
	return result
}

// Multiply takes two numbers and returns the
// result of multiplying one by another
func Multiply(inputs ...float64) float64 {
	var result float64 = 1
	for _, input := range inputs {
		result *= input
	}
	return result
}

// Divide takes two numbers and returns the
// result of dividing one by another
func Divide(inputs ...float64) (float64, error) {
	var result float64 = 1
	for i, input := range inputs {
		if i == 0 {
			result = input
		} else {
			if input == 0 {
				return 0, fmt.Errorf("bad input: %v, %f (division by zero is undefined)", inputs, input)
			}
			result /= input
		}
	}
	return result, nil
}

// Sqrt takes a positive number and returns its square root
func Sqrt(a float64) (float64, error) {
	if a > 0 {
		return math.Sqrt(a), nil
	}
	return 0, fmt.Errorf("bad input: %f sqrt of negative numbers is not allowed", a)
}

// Evaluate receives a string with an aritmetic operation and returns the result
// only expressions with a floating point value followed by one or more spaces
// followed by an aritmentic operator *,+,/,- followed by one or more spaces
// followed by a floating point value are accepted.
func Evaluate(expr string) (float64, error) {
	var a float64
	var b float64
	var op string
	n, err := fmt.Sscanf(expr, "%f%s%f\n", &a, &op, &b)

	if err != nil {
		return 0, fmt.Errorf("%s Unexpected error %s", expr, err)
	}
	if n > 3 {
		return 0, fmt.Errorf("%s Invalid expression. Wrong number or argument %d", expr, n)
	}

	switch operation := op; operation {
	case "+":
		return Add(a, b), nil
	case "-":
		return Substract(a, b), nil
	case "*":
		return Multiply(a, b), nil
	case "/":
		result, err := Divide(a, b)
		if err != nil {
			return 0, err
		}
		return result, nil
	default:
		return 0, fmt.Errorf("%s Invalid operator %s", expr, op)
	}
}
