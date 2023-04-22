package main

type calculatorFunctionsProvider interface {
	add() int 
	sub() int
	div() int 
	mult() int 
}

type Values struct {
	a int
	b int
}

// add is a calculatorFunctions that receives two ints and returns their sum.
func (cf calculatorFunctionsProvider)add(a, b int) int {
	return a + b
}

func (cf calculatorFunctionsProvider)sub(a, b int) int {
	return a - b
}

func (cf calculatorFunctionsProvider)div(a, b int) int {
	if b == 0 {
		panic("division by zero is undefined")
	} else if b > a {
		panic("numerator smaller than denominator")
	}
	return a / b
}

func (cf calculatorFunctionsProvider)mult(a, b int) int {
	return a * b
}

func main() {
	
}