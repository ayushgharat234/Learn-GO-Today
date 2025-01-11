// Package main demonstrates the use of data types, variables, constants, type inference, and default values in Go.
package main

import "fmt"

func main() {
	// SECTION 1: Variables with Explicit Types
	// Declaring variables with explicit types ensures clarity about the variable's purpose and data type.
	var name string = "Alice"  // A string variable explicitly initialized to "Alice"
	var age int = 30           // An integer variable explicitly initialized to 30
	var height float64 = 5.9   // A float64 variable explicitly initialized to 5.9
	var isStudent bool = false // A boolean variable explicitly initialized to false

	fmt.Println("SECTION 1: Explicit Type Variables")
	fmt.Printf("Name: %s, Age: %d, Height: %.1f, Is Student: %t\n\n", name, age, height, isStudent)

	// SECTION 2: Type Inference
	// Go can automatically infer the type based on the value assigned during initialization.
	inferredName := "Bob"   // Inferred as a string
	inferredAge := 25       // Inferred as an int
	inferredHeight := 6.2   // Inferred as a float64
	inferredStudent := true // Inferred as a bool

	fmt.Println("SECTION 2: Type Inference")
	fmt.Printf("Name: %s, Age: %d, Height: %.1f, Is Student: %t\n\n", inferredName, inferredAge, inferredHeight, inferredStudent)

	// SECTION 3: Default Values (Zero Values)
	// Variables declared without initialization are assigned default "zero" values.
	var zeroString string // Default value is an empty string ""
	var zeroInt int       // Default value is 0
	var zeroFloat float64 // Default value is 0.0
	var zeroBool bool     // Default value is false

	fmt.Println("SECTION 3: Default (Zero) Values")
	fmt.Printf("String: '%s', Int: %d, Float: %.1f, Bool: %t\n\n", zeroString, zeroInt, zeroFloat, zeroBool)

	// SECTION 4: Constants
	// Constants are immutable and must be initialized at the time of declaration.
	const pi float64 = 3.14159 // A float64 constant representing the value of π
	const gravity = 9.8        // Type inference also works for constants (inferred as float64)

	fmt.Println("SECTION 4: Constants")
	fmt.Printf("Pi: %.5f, Gravity: %.1f\n\n", pi, gravity)

	// SECTION 5: Data Types in Go
	// Demonstrating various data types and their usage.

	// Integer types
	var smallInt int8 = -128                   // 8-bit signed integer: Range -128 to 127
	var largeInt uint64 = 18446744073709551615 // 64-bit unsigned integer: 0 to 2^64-1

	// Floating-point types
	var singlePrecision float32 = 3.14         // 32-bit floating-point
	var doublePrecision float64 = 3.1415926535 // 64-bit floating-point

	// Complex numbers
	var complexNum complex128 = complex(5, 2) // Complex numbers: (5 + 2i)

	// Rune and Byte
	var aRune rune = 'A' // Runes are Unicode code points (alias for int32)
	var aByte byte = 'B' // Byte is an alias for uint8

	fmt.Println("SECTION 5: Data Types")
	fmt.Printf("Small Int: %d, Large Int: %d\n", smallInt, largeInt)
	fmt.Printf("Float32: %.2f, Float64: %.10f\n", singlePrecision, doublePrecision)
	fmt.Printf("Complex: %.1f + %.1fi\n", real(complexNum), imag(complexNum))
	fmt.Printf("Rune: %c (Unicode: %U), Byte: %c\n\n", aRune, aRune, aByte)

	// SECTION 6: Conversion between Data Types
	// Explicit type conversion is required as Go is strongly typed.
	intVal := 42
	floatVal := float64(intVal) // Converting int to float64
	uintVal := uint(floatVal)   // Converting float64 to uint

	fmt.Println("SECTION 6: Type Conversion")
	fmt.Printf("Int: %d, Float64: %.2f, Uint: %d\n\n", intVal, floatVal, uintVal)

	// SECTION 7: Working with Booleans
	isAdult := age >= 18 // Boolean expressions
	fmt.Println("SECTION 7: Booleans")
	fmt.Printf("Is Adult (Age >= 18): %t\n", isAdult)

	// SECTION 8: String Operations
	greeting := "Hello"
	audience := "World"
	combined := greeting + ", " + audience + "!" // String concatenation
	fmt.Println("SECTION 8: Strings")
	fmt.Printf("Combined String: %s\n", combined)
}
