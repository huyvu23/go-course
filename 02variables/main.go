package main

import "fmt"

// this is a constant variable
// Constants are immutable values that are known at compile time and do not change for the life of the program.
// They are declared using the const keyword and can be of any type.

const LoginToken string = "abc123"

func main() {
	var a int = 1
	var b int = 2
	var c int = a + b
	var name string = "Huy"
	fmt.Println(name)
	fmt.Printf("Type of variables a: %T\n", a)
	fmt.Println(c)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of variables a: %T\n", isLoggedIn)

	// explain the code uint8
	// uint8 is an unsigned integer type that can hold values from 0 to 255.
	// It is commonly used when you need to store small numbers or when memory efficiency is a concern.
	// In this code, we declare a variable smallVal of type uint8 and assign it the value 255.
	// This means that smallVal can hold any value from 0 to 255, and it is not allowed to hold negative values.
	var smallVal uint8 = 255
	fmt.Printf("Type of variables a: %T\n", smallVal)

	var smallFloat float32 = 255.3
	fmt.Printf("Type of variables a: %T\n", smallFloat)

	//default value of variables
	var defaultVal int
	fmt.Printf("Type of variables a: %T\n", defaultVal)
	// defaultVal is an integer variable that is declared but not initialized.
	// In Go, when a variable is declared without an explicit initialization, it is assigned a default value.
	// For numeric types like int, the default value is 0.
	// This means that defaultVal will have a value of 0 until it is explicitly assigned a different value.
	// The zero value is the default value for a variable of a specific type in Go.

	// implicitly typed variables
	var x = 10
	fmt.Printf("Type of variables x: %T\n", x)

	// no var style
	//Can only be used inside functions
	// Variable declaration and value assignment cannot be done separately (must be done in the same line)
	age := 20
	fmt.Printf("Type of variables age: %T\n", age)
}
