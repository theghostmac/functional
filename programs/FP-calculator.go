package main

import (
	"fmt"
)

// calculatorFunctions is a type aliased function.
type calculatorFunctions func(int, int) int

// add is a calculatorFunctions that receives two ints and returns their sum.
func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func div(a, b int) int {
	if b == 0 {
		panic("division by zero is undefined")
	} else if b > a {
		panic("numerator smaller than denominator")
	}
	return a / b
}

func mult(a, b int) int {
	return a * b
}

var (
	// operations are functions possible on calculatorFunctions types, mapped as strings.
	operations = map[string]calculatorFunctions{
		"+": add,
		"-": sub,
		"/": div,
		"*": mult,
	}
)

// caluculate selects the operation as a string, and acts it on two integers | map dispatcher.
func caluculate(a, b int, selectedOperation string) int {
	// if the selectedOperation matches a key in the map, perform the operation's function.
	if operation, ok := operations[selectedOperation]; ok {
		return operation(a, b)
	}
	// panic if the selectedOperation does not match a key in the map.
	panic("operation is not defined")
}

func main() {
	fpTestedCalculator := caluculate(10, 5, "+")
	fmt.Println(fpTestedCalculator)
}