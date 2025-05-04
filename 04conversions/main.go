package main

import (
	"fmt"
	"strconv"
)

func main() {
	// This is a simple Go program that demonstrates the use of type conversion.
	// It converts an integer to a float64 and prints the result.
	// The program also shows how to convert a string to an integer using the strconv package.
	// The main function is the entry point of the program, and the fmt package is used for formatted I/O.
	// The Println function is used to print the string to the console, followed by a newline character.

	var num int = 42
	var floatNum float64 = float64(num)
	fmt.Println("Converted integer to float64:", floatNum)

	var age int = 20
	// Convert int to string using strconv.Itoa
	// The strconv package provides functions for converting between strings and other data types.
	// The Itoa function converts an integer to a string.
	// The result is a string representation of the integer.
	convertAge := strconv.Itoa(age)
	fmt.Println("Converted integer to string:", convertAge)

	var starRate = "4"
	// Convert string to int using strconv.Atoi
	// The Atoi function converts a string to an integer.
	// It returns the converted value and an error if the conversion fails.
	// The result is an integer representation of the string.
	convertStarRate, err := strconv.Atoi(starRate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Converted string to integer:", convertStarRate)
}
