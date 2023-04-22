package main

import (
	"fmt"
)

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

func calculate(a, b int, operation string) int {
	switch operation {
	case "+":
		return add(a, b)
	case "-":
		return sub(a, b)
	case "*":
		return mult(a, b)
	case "/":
		return div(a, b)
	default:
		panic("operation not defined initially")
	}
}

func main() {
	testResult := calculate(10, 5, "+")
	fmt.Println(testResult)
}