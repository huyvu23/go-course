package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

// A struct (short for structure) is used to create a collection of members of different data types, into a single variable.
func main() {

	// Access Struct Members
	var pers1 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	fmt.Println("Person1:", pers1)

	// Print Pers1 info by calling a function
	printPerson(pers1)
}

// Pass Struct as Function Arguments
func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
