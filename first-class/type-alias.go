package main

import (
	"fmt"
)

/* type aliases:
 * improves communication intent.
 * helps error messages sound clearer.
 * allows you to attach functions to them (like age below).
 */

type (
	// phoneNumber is a string type alias with clearer intent.
	phoneNumber string
	// age is a uint type alias with clearer intent.
	age uint
)

// Person is a struct model with name and phone number of a person.
type Person struct {
	Name string
	age age
	phone phoneNumber
}

// setPhoneNumber receives a phoneNumber belonging to a person.
func (p *Person) setPhoneNumber(setPhone phoneNumber) {
	p.phone = setPhone
}

// valid is bound to the age type.
func (a age) valid() bool {
	return a < 120
}

// isValidPerson checks if a Person's age and name is valid.
func isValidPerson(p Person) bool {
	return p.age.valid() && p.Name != ""
}

func main() {
	// Mac is a programmer implementing Person struct.
	Mac := Person{
		Name: "GhostMac",
		phone: "1234567890",
	}

	fmt.Println("%v\n", Mac)
}