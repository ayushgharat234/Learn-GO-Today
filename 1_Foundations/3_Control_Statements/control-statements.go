// Package main demonstrates conditional statements (if-else), loops (for), and switch-case constructs in Go.
package main

import "fmt"

func main() {
	// SECTION 1: Conditional Statements (if-else)
	fmt.Println("SECTION 1: Conditional Statements")

	age := 17 // A sample variable for age
	if age > 18 {
		// This block executes if the condition (age > 18) is true
		fmt.Println("You are an Adult.")
	} else if age == 18 {
		// This block executes if the condition (age == 18) is true
		fmt.Println("You just turned 18!")
	} else {
		// This block executes if none of the above conditions are true
		fmt.Println("You are a Minor.")
	}

	// Using a boolean expression directly in the if condition
	isWeekend := false
	if isWeekend {
		fmt.Println("Relax, it's the weekend!")
	} else {
		fmt.Println("Time to work!")
	}
	fmt.Println() // Adding a blank line for output readability

	// SECTION 2: Loops in Go
	fmt.Println("SECTION 2: Loops")

	// Basic for loop
	fmt.Println("Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}
	fmt.Println()

	// For loop as a while loop
	fmt.Println("Using 'for' as a while loop:")
	counter := 3
	for counter > 0 {
		fmt.Printf("Counter: %d\n", counter)
		counter-- // Decrement the counter
	}
	fmt.Println()

	// Infinite loop with a break statement
	fmt.Println("Infinite loop with break:")
	count := 0
	for {
		if count == 3 {
			fmt.Println("Breaking out of the loop!")
			break // Exits the loop when count == 3
		}
		fmt.Printf("Count: %d\n", count)
		count++
	}
	fmt.Println()

	// Using continue to skip an iteration
	fmt.Println("Loop with continue statement:")
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue // Skips the iteration for even numbers
		}
		fmt.Printf("Odd Number: %d\n", i)
	}
	fmt.Println()

	// SECTION 3: Switch Statement
	fmt.Println("SECTION 3: Switch Statement")

	day := 3 // A variable representing the day of the week
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4, 5:
		// Grouping multiple cases together
		fmt.Println("Thursday or Friday")
	default:
		// Default case for unmatched conditions
		fmt.Println("Weekend")
	}

	// Switch with no condition (acting as an if-else ladder)
	fmt.Println("Switch with no condition:")
	ageCategory := 25
	switch {
	case ageCategory < 18:
		fmt.Println("Underage")
	case ageCategory >= 18 && ageCategory < 60:
		fmt.Println("Working age")
	default:
		fmt.Println("Senior citizen")
	}
	fmt.Println()

	// Fallthrough behavior
	fmt.Println("Switch with fallthrough:")
	switch level := 2; level {
	case 1:
		fmt.Println("Level 1")
		fallthrough // Forces execution of the next case
	case 2:
		fmt.Println("Level 2")
	case 3:
		fmt.Println("Level 3")
	default:
		fmt.Println("Unknown level")
	}
	fmt.Println()

	// Switch with a variable declaration in the statement
	fmt.Println("Switch with variable declaration:")
	switch num := 15; {
	case num%2 == 0:
		fmt.Println("Even number")
	case num%2 != 0:
		fmt.Println("Odd number")
	}
}
