// Package main demonstrates error handling in Go, including basic error values, error creation, wrapping, sentinel errors, custom error types, and best practices.
package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// SECTION 1: Basic Error Handling
	fmt.Println("SECTION 1: Basic Error Handling")
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	fmt.Println()

	// SECTION 2: Creating and Wrapping Errors
	fmt.Println("SECTION 2: Creating and Wrapping Errors")
	err = readConfig()
	if err != nil {
		wrappedErr := fmt.Errorf("loadApp failed: %w", err)
		fmt.Println("Wrapped Error:", wrappedErr)
	}
	fmt.Println()

	// SECTION 3: Sentinel Errors
	fmt.Println("SECTION 3: Sentinel Errors")
	err = findUser(42)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("User not found (sentinel error)")
	}
	fmt.Println()

	// SECTION 4: Custom Error Types
	fmt.Println("SECTION 4: Custom Error Types")
	err = fetch()
	if httpErr, ok := err.(*HTTPError); ok {
		fmt.Printf("Custom Error - Status code: %d, Message: %s\n", httpErr.Code, httpErr.Message)
	}
	fmt.Println()

	// SECTION 5: Anti-Patterns (for demonstration only)
	fmt.Println("SECTION 5: Anti-Patterns")
	// Don't do this: ignoring errors
	_, err = divide(1, 0)
	// _ = err // BAD: error ignored
	if err != nil {
		fmt.Println("Handled error instead of ignoring:", err)
	}
	fmt.Println()

	// SECTION 6: Real-World Example: Error Propagation
	fmt.Println("SECTION 6: Real-World Example: Error Propagation")
	if err := process(); err != nil {
		log.Fatalf("process failed: %v", err)
	}
}

// divide returns the result of a / b, or an error if b is zero.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// readConfig simulates a function that returns a simple error.
func readConfig() error {
	return errors.New("config file not found")
}

// Sentinel error for demonstration.
var ErrNotFound = errors.New("not found")

// findUser simulates a function that returns a sentinel error.
func findUser(id int) error {
	return ErrNotFound
}

// HTTPError is a custom error type with additional fields.
type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
}

// fetch returns a custom error type.
func fetch() error {
	return &HTTPError{Code: 404, Message: "Not Found"}
}

// Metadata is a placeholder struct for the real-world example.
type Metadata struct {
	ID   int
	Name string
}

// fetchMetadata simulates fetching data and returns an error.
func fetchMetadata() (Metadata, error) {
	return Metadata{}, errors.New("failed to fetch metadata")
}

// saveToDB simulates saving data and returns an error.
func saveToDB(data Metadata) error {
	return nil // Simulate success
}

// process demonstrates error propagation with context.
func process() error {
	data, err := fetchMetadata()
	if err != nil {
		return fmt.Errorf("fetch failed: %w", err)
	}
	if err := saveToDB(data); err != nil {
		return fmt.Errorf("db save failed: %w", err)
	}
	return nil
} 