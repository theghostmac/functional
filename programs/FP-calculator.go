package main

type calculatorFunctions func(int, int) int

var (
	operations = map[string]calculatorFunctions{
		"+": add,
		"-": sub,
		"/": div,
		"*": mult,
	}
)

func caluculate(a, b int, opString string)

func main() {
	
}