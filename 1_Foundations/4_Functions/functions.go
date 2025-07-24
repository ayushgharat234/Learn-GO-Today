// Package main demonstrates the use of functions in Go, including basic, parameterized,
// variadic, anonymous, closures, recursion, and function types.
package main

import "fmt"

// SECTION 1: Basic Function
// greet returns a greeting message with the provided name.
func greet(name string) string {
	return "Hello, " + name
}

// SECTION 2: Function with Multiple Return Values
// calculate takes two integers and returns their sum and difference.
func calculate(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// SECTION 3: Variadic Function
// sumAll calculates the sum of an arbitrary number of integers.
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// SECTION 4: Anonymous Function
// Demonstrates how to use an unnamed function.
func demonstrateAnonymousFunction() {
	anonymous := func(a, b int) int {
		return a * b
	}
	fmt.Printf("Multiplication of 3 and 4: %d\n", anonymous(3, 4))
}

// SECTION 5: Closure
// showClosure demonstrates a function that captures and uses an external variable.
func showClosure() func() {
	counter := 0
	return func() {
		counter++
		fmt.Printf("Counter: %d\n", counter)
	}
}

// SECTION 6: Recursive Function
// factorial calculates the factorial of a number using recursion.
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// SECTION 7: Function Types
// type operation defines a function type that takes two integers and returns an integer.
type operation func(int, int) int

// add and multiply implement the operation function type.
func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

// demonstrateFunctionType shows how to use a function type as a parameter.
func demonstrateFunctionType(op operation, a, b int) {
	fmt.Printf("Result of operation: %d\n", op(a, b))
}

// SECTION 8: Higher-Order Functions
// higherOrder takes a function as a parameter and applies it to two integers.
func higherOrder(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

// main is the entry point for the program.
func main() {
	// SECTION 1: Calling a Basic Function
	fmt.Println("SECTION 1: Basic Function")
	message := greet("Alice")
	fmt.Println(message)
	// Practice: Try calling greet() with your own name.
	fmt.Println()

	// SECTION 2: Using a Function with Multiple Return Values
	fmt.Println("SECTION 2: Function with Multiple Return Values")
	sum, diff := calculate(10, 5)
	fmt.Printf("Sum: %d, Difference: %d\n", sum, diff)
	// Practice: Create a new function that returns product and quotient of two numbers.
	fmt.Println()

	// SECTION 3: Variadic Function
	fmt.Println("SECTION 3: Variadic Function")
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Printf("Sum of numbers: %d\n", total)
	// Practice: Modify sumAll to return the average as well.
	fmt.Println()

	// SECTION 4: Anonymous Function
	fmt.Println("SECTION 4: Anonymous Function")
	demonstrateAnonymousFunction()
	// Practice: Write an anonymous function that returns square of a number.
	fmt.Println()

	// SECTION 5: Closure
	fmt.Println("SECTION 5: Closure")
	increment := showClosure()
	increment()
	increment()
	// Practice: Try creating a closure that accumulates sum.
	fmt.Println()

	// SECTION 6: Recursive Function
	fmt.Println("SECTION 6: Recursive Function")
	fmt.Printf("Factorial of 5: %d\n", factorial(5))
	// Practice: Write a recursive function to compute Fibonacci numbers.
	fmt.Println()

	// SECTION 7: Function Types
	fmt.Println("SECTION 7: Function Types")
	demonstrateFunctionType(add, 10, 5)
	demonstrateFunctionType(multiply, 10, 5)
	// Practice: Create a new operation type function that divides two numbers.
	fmt.Println()

	// SECTION 8: Higher-Order Functions
	fmt.Println("SECTION 8: Higher-Order Functions")
	result := higherOrder(3, 4, func(x, y int) int {
		return x * y
	})
	fmt.Printf("Result of higher-order function: %d\n", result)
	// Practice: Use higherOrder with subtraction logic.
}
