package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome user input"
	fmt.Println(welcome)

	// This code snippet is reading user input from the console. Here's a breakdown of what each line is
	// doing:
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	// The line `name, _ := reader.ReadString('\n')` is reading the user input from the console until a
	// newline character ('\n') is encountered.
	// _ is used to ignore the second return value, which is an error.
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello", name)
}
