package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	john := person{
		firstName: "John",
		lastName:  "Reyes",
		contact: contactInfo{
			email:   "john@yopmail.com",
			zipCode: 94000,
		},
	}
	fmt.Printf("%+v", john)
}
