package main

import (
	"github.com/golang-collections/collections/stack"
)

var calc Calculator

func main() {}

// Operator is used to specify calculator operator.
type Operator int

const (
	// UNKNOWN operator is used for unknown operators
	UNKNOWN Operator = iota
	// ADD operator is used to add two numbers
	ADD
	// SUBTRACT operator is used to subtract two numbers
	SUBTRACT
	// MULTIPLY operator is used to multiply two numbers
	MULTIPLY
	// DIVIDE operator is used to divide two numbers
	DIVIDE
)

func (o Operator) String() string {
	switch o {
	case ADD:
		return "add"
	case SUBTRACT:
		return "subtract"
	case MULTIPLY:
		return "multiply"
	case DIVIDE:
		return "divide"
	default:
		return ""
	}
}

// ParseOperator is used to convert string into Calculator Operator
func ParseOperator(operator string) Operator {
	switch operator {
	case "add":
		return ADD
	case "subtract":
		return SUBTRACT
	case "multiply":
		return MULTIPLY
	case "divide":
		return DIVIDE
	default:
		return UNKNOWN
	}
}

// Calculator object performs calculations.
type Calculator struct {
	numbers stack.Stack
	result  int
}

func (c *Calculator) enterNumber(n int) {
	c.numbers.Push(n)
}

func (c *Calculator) press(operator Operator) {
	n2 := c.numbers.Pop().(int)
	n1 := c.numbers.Pop().(int)

	switch operator {
	case ADD:
		c.result = n2 + n1
	case SUBTRACT:
		c.result = n1 - n2
	case MULTIPLY:
		c.result = n1 * n2
	case DIVIDE:
		c.result = n1 / n2
	}
}
