package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var john person

	john.firstName = "John"
	john.lastName = "Reyes"

	fmt.Println(john)
	fmt.Printf("%+v", john)
}
