package main

import (
	"fmt"
)

type (
	// boolReturner is a funtion that returns bool. functins are firs class citizens.
	boolReturner func(int) bool
)

// largerThanTwo is a simple function that returns any value greater than 2.
func largerThanTwo(i int) bool {
	return i > 2
}

// filter receives an integer array and a boolean aliased function and returns another array.  
func filter(integers []int, bR boolReturner)  []int {
	set := []int{}
	fmt.Println(set)
	for _, values := range integers {
		if bR(values) {
			set = append(set, values)
		}
	}
	return set
}

func main() {
	ints := []int{1, 2, 3}
	
	fmt.Println(filter(ints, largerThanTwo))

	// in-line functions are functions in variables.
	inlineFunction := func(i int) bool {return i > 2}
	filter(ints, inlineFunction)

	// anonymous functions have no names
	filter([]int{1,2,3}, func(i int) bool {return i > 2})
}