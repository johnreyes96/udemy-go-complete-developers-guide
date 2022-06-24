package main

import "testing"

func TestUpdateName(t *testing.T) {
	john := person{
		firstName: "John",
		lastName:  "Reyes",
		contactInfo: contactInfo{
			email:   "john@yopmail.com",
			zipCode: 94000,
		},
	}

	john.updateName("Johny")

	if john.firstName != "Johny" {
		t.Errorf("Expected Johny, but got %v", john.firstName)
	}
}
