package main

import (
	"fmt"
)

// create function as a stand-alone type.
type predicate func(int) bool

type (
	// boolReturner is a funtion that returns bool. functins are firs class citizens.
	boolReturner func(int) bool
)

// create function to return a function.
func createLargerThanPredicate(threshold int) predicate {
	return func(i int) bool {
		return i > threshold
	}
}

// functions as struct field values.
type ConstraintChecker struct {
	largerThan predicate
	smallerThan predicate
}

// create function that takes another function as parameter.
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

var (
	largerThanTwo = createLargerThanPredicate(2)
	largerThanFive = createLargerThanPredicate(5)
	largerThanTen = createLargerThanPredicate(10)
)

func main() {
	ints := []int{1,2,3}
	
	// create function as variable value.
	largerThanTwo := createLargerThanPredicate(2)
	filter(ints, largerThanTwo)
	
	// create inline functions.
	inlineFunction := func(i int) bool {return i > 2}
	filter(ints, inlineFunction)

	// anonymous functions have no names.
	filter([]int{1,2,3}, func(i int) bool {return i > 2})
	
	predicates := predicate{
		largerThanTwo,
		largerThanFive,
		largerThanTen,
	}
	for _, predicate := range predicates {
		fmt.Printf("%v\n", filter(ints, predicate))
	}
}