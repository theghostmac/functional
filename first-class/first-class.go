package main

import (
	"fmt"
)

// type aliases.
// improves communication intent.
// helps error messages sound clearer.

type (
	// phoneNumber is a string type alias with clearer intent.
	phoneNumber string
)

// Person is a struct model with name and phone number of a person.
type Person struct {
	Name string
	phone phoneNumber
}

// setPhoneNumber receives a phoneNumber belonging to a person.
func (p *Person) setPhoneNumber(setPhone phoneNumber) {
	p.phone = setPhone
}

func main() {
	// Mac is a person implementing Person struct.
	Mac := Person{
		Name: "GhostMac",
		phone: "1234567890",
	}
	
	fmt.Println("%v\n", Mac)
}