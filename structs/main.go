package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	john := person{
		firstName: "John",
		lastName:  "Reyes",
		contactInfo: contactInfo{
			email:   "john@yopmail.com",
			zipCode: 94000,
		},
	}

	john.updateName("Johny")
	john.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
