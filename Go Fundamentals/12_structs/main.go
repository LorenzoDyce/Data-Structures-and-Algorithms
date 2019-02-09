package main

import (
	"fmt"
	"strconv"
)

// Define "Person struct"
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting Method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " +
		strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastNAme string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastNAme
	}
}

func main() {
	// Init person using struct
	person1 := Person{
		firstName: "Sam",
		lastName:  "Smith",
		city:      "Toronto",
		gender:    "f",
		age:       25}

	person2 := Person{
		"Bob",
		"Johnson",
		"Mississauga",
		"m",
		30}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)

	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Thompson")
	fmt.Println(person2.greet())

}
