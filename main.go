package main

import "fmt"

func main() {

	// Declare variables
	var name string
	var age int

	// Ask user to enter name
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	// &name means: store the entered value inside variable "name"

	// Ask user to enter age
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	// Print using formatting
	fmt.Printf("Hello %s, you are %d years old.\n", name, age)

	// If-Else Condition
	if age >= 18 {
		fmt.Println("You are eligible to vote ✅")
	} else {
		fmt.Println("You are not eligible to vote ❌")
	}
}
