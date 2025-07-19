// Package main demonstrates basic variable declarations, assignments, and printing in Go.
package main

import (
	"fmt"
)

// main is the entry point of the program.
// It showcases different ways to declare, assign, and print variables of various types in Go.
func main() {
	// Declare a string variable and assign a value
	var name string
	name = "Hemant"
	fmt.Println("name: ", name)

	// Declare an integer variable and assign a value
	var age int
	age = 20
	fmt.Println("age: ", age)

	// Declare a boolean variable and assign a value
	var isStudent bool
	isStudent = true
	fmt.Println("isStudent: ", isStudent)

	// Short variable declaration for a string
	company := "Meta"
	fmt.Println("company: ", company)

	// Declare and initialize multiple string variables in a block
	var (
		firstName = "Hemant"
		lastName  = "Kumar"
	)
	fmt.Println("firstName: ", firstName, "\t", "lastName: ", lastName)

	// Short variable declaration for multiple variables of different types
	car, cost := "Audi", 5000.50

	fmt.Printf("Care %s Cost %f \n", car, cost)

	// Declare multiple float64 variables in a block
	var (
		salary float64
		bonus  float64
	)

	// Assign values to the float64 variables
	salary = 10000.50
	bonus = 1000.50

	fmt.Println("salary: ", salary, "bonus: ", bonus)

	// Declare multiple integer variables with zero values
	var i, j int
	fmt.Println("i: ", i, "j: ", j)
}
