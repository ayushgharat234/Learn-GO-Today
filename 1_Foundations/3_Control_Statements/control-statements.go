// Package main demonstrates conditional statements (if-else), loops (for), and switch-case constructs in Go.
package main

import "fmt"

func main() {
	// SECTION 1: Conditional Statements (if-else)
	fmt.Println("SECTION 1: Conditional Statements")

	age := 17 // A sample variable for age
	if age > 18 {
		fmt.Println("You are an Adult.")
	} else if age == 18 {
		fmt.Println("You just turned 18!")
	} else {
		fmt.Println("You are a Minor.")
	}

	// Real-world example: Access control
	userLoggedIn := true
	userRole := "admin"
	if userLoggedIn && userRole == "admin" {
		fmt.Println("Welcome Admin! Access granted.")
	} else {
		fmt.Println("Access Denied.")
	}

	isWeekend := false
	if isWeekend {
		fmt.Println("Relax, it's the weekend!")
	} else {
		fmt.Println("Time to work!")
	}
	fmt.Println()

	// SECTION 2: Loops in Go
	fmt.Println("SECTION 2: Loops")

	fmt.Println("Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}
	fmt.Println()

	fmt.Println("Using 'for' as a while loop:")
	counter := 3
	for counter > 0 {
		fmt.Printf("Counter: %d\n", counter)
		counter--
	}
	fmt.Println()

	fmt.Println("Infinite loop with break:")
	count := 0
	for {
		if count == 3 {
			fmt.Println("Breaking out of the loop!")
			break
		}
		fmt.Printf("Count: %d\n", count)
		count++
	}
	fmt.Println()

	fmt.Println("Loop with continue statement:")
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("Odd Number: %d\n", i)
	}
	fmt.Println()

	// Real-world loop example: Sum of first N numbers
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("Sum of 1 to 100: %d\n\n", sum)

	// SECTION 3: Switch Statement
	fmt.Println("SECTION 3: Switch Statement")

	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4, 5:
		fmt.Println("Thursday or Friday")
	default:
		fmt.Println("Weekend")
	}

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

	fmt.Println("Switch with fallthrough:")
	switch level := 2; level {
	case 1:
		fmt.Println("Level 1")
		fallthrough
	case 2:
		fmt.Println("Level 2")
	case 3:
		fmt.Println("Level 3")
	default:
		fmt.Println("Unknown level")
	}
	fmt.Println()

	fmt.Println("Switch with variable declaration:")
	switch num := 15; {
	case num%2 == 0:
		fmt.Println("Even number")
	case num%2 != 0:
		fmt.Println("Odd number")
	}
}
