package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   string
}

func main() {
	contact := contactInfo{email: "alex.anderson@smth.com", zip: "98059"}
	alex := person{firstName: "Alex", lastName: "Anderson", contact: contact}

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email: "jim.party@gmail.com",
			zip:   "98059",
		},
	}

	alex.print()
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
