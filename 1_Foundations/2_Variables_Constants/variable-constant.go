// Package main demonstrates the use of data types, variables, constants,
// type inference, and default values in Go with practical examples.
package main

import "fmt"

func main() {
	// SECTION 1: Variables with Explicit Types
	// Declaring variables with explicit types ensures clarity about the variable's purpose and data type.
	var name string = "Alice"
	var age int = 30
	var height float64 = 5.9
	var isStudent bool = false

	fmt.Println("SECTION 1: Explicit Type Variables")
	fmt.Printf("Name: %s, Age: %d, Height: %.1f, Is Student: %t\n\n", name, age, height, isStudent)
	// Example Use Case: Defining user profile fields explicitly in a config.

	// SECTION 2: Type Inference
	inferredName := "Bob"
	inferredAge := 25
	inferredHeight := 6.2
	inferredStudent := true

	fmt.Println("SECTION 2: Type Inference")
	fmt.Printf("Name: %s, Age: %d, Height: %.1f, Is Student: %t\n\n", inferredName, inferredAge, inferredHeight, inferredStudent)
	// Example Use Case: Local function variables inferred based on returned data.

	// SECTION 3: Default Values (Zero Values)
	var zeroString string
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool

	fmt.Println("SECTION 3: Default (Zero) Values")
	fmt.Printf("String: '%s', Int: %d, Float: %.1f, Bool: %t\n\n", zeroString, zeroInt, zeroFloat, zeroBool)
	// Example Use Case: Default values in HTTP struct configurations.

	// SECTION 4: Constants
	const pi float64 = 3.14159
	const gravity = 9.8

	fmt.Println("SECTION 4: Constants")
	fmt.Printf("Pi: %.5f, Gravity: %.1f\n\n", pi, gravity)

	// Constant block using iota
	const (
		_ = iota
		Low
		Medium
		High
	)
	fmt.Println("SECTION 4B: Constants with iota")
	fmt.Printf("Low: %d, Medium: %d, High: %d\n\n", Low, Medium, High)
	// Example Use Case: Priority levels in task management systems.

	// SECTION 5: Data Types in Go
	var smallInt int8 = -128
	var largeInt uint64 = 18446744073709551615
	var singlePrecision float32 = 3.14
	var doublePrecision float64 = 3.1415926535
	var complexNum complex128 = complex(5, 2)
	var aRune rune = 'A'
	var aByte byte = 'B'

	fmt.Println("SECTION 5: Data Types")
	fmt.Printf("Small Int: %d, Large Int: %d\n", smallInt, largeInt)
	fmt.Printf("Float32: %.2f, Float64: %.10f\n", singlePrecision, doublePrecision)
	fmt.Printf("Complex: %.1f + %.1fi\n", real(complexNum), imag(complexNum))
	fmt.Printf("Rune: %c (Unicode: %U), Byte: %c\n\n", aRune, aRune, aByte)

	emoji := "🚀"
	runes := []rune(emoji)
	fmt.Printf("Rune Count in '%s': %d, Unicode: %U\n", emoji, len(runes), runes[0])
	// Example Use Case: Logging emoji status in APIs or converting UTF-8 streams.

	// SECTION 6: Conversion between Data Types
	intVal := 42
	floatVal := float64(intVal)
	uintVal := uint(floatVal)

	fmt.Println("SECTION 6: Type Conversion")
	fmt.Printf("Int: %d, Float64: %.2f, Uint: %d\n\n", intVal, floatVal, uintVal)
	// Be cautious: Converting float64 to uint truncates the decimal part.
	// Example Use Case: Parsing JSON numbers or converting between types in API handlers.

	// SECTION 7: Working with Booleans
	isAdult := age >= 18
	fmt.Println("SECTION 7: Booleans")
	fmt.Printf("Is Adult (Age >= 18): %t\n", isAdult)
	// Example Use Case: Authorization logic based on user properties.

	// SECTION 8: String Operations
	greeting := "Hello"
	audience := "World"
	combined := greeting + ", " + audience + "!"
	fmt.Println("SECTION 8: Strings")
	fmt.Printf("Combined String: %s\n", combined)

	config := `{
	"env": "prod",
	"debug": false
	}`
	fmt.Println("Raw Config:\n", config)
	// Example Use Case: Generating raw config templates for cloud deployment.

	// SECTION 9: Best Practices Summary
	// - Use ':=' for short-lived local variables.
	// - Use 'var' for zero-values or explicit declarations.
	// - Use 'const' for values that never change.
	// - Always convert types explicitly to avoid silent bugs.
	// - Understand zero values — they're idiomatic in Go.
}
