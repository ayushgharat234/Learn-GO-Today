// Package main demonstrates a simple Go program that prints a greeting message to the console.
package main

// Importing the "fmt" package, which provides functions for formatted I/O operations.
// The fmt package is used to print output, read input, and format strings.
import (
	"fmt"
)

// The main function is the entry point of the Go program.
// Every Go program must have a main function in the "main" package to execute.
func main() {
	// Declaring and initializing a variable 'name' with the value "World".
	// The ':=' syntax is used to declare a variable with automatic type inference.
	// Here, 'name' will be of type string because "World" is a string literal.
	// Go automatically infers the variable type based on the value assigned.
	name := "World"

	// The fmt.Printf function is used to print formatted output to the console.
	// The "%s" format specifier is used to insert the value of 'name' (a string) into the output string.
	// In this case, it prints the message: "Hello, World!"
	// The '\n' at the end ensures the cursor moves to the next line after printing.
	// fmt.Printf allows for more advanced formatting options if needed.
	fmt.Printf("Hello, %s!\n", name)

	// Notes on fmt.Printf:
	// - "%s" is used to print a string.
	// - "\n" represents a newline character, ensuring the output appears cleanly on its own line.
	// - The '%s' format specifier is useful when inserting variables like strings into a formatted output.
}
