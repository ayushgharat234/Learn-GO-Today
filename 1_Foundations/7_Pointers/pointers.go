// Package main demonstrates pointers in Go,
// along with their key operations, properties, and use cases.
package main

import "fmt"

func main() {
	// SECTION 1: Pointer Basics
	fmt.Println("SECTION 1: Pointer Basics")
	// Declaring and initializing a pointer.
	var x int = 42
	var ptr *int = &x // Pointer to x
	fmt.Printf("Value of x: %d, Address of x: %p\n", x, &x)
	fmt.Printf("Value of ptr: %p, Value at ptr: %d\n", ptr, *ptr)

	// Modifying the value of x through the pointer.
	*ptr = 100
	fmt.Println("Updated value of x:", x)
	fmt.Println()

	// SECTION 2: Pointers and Functions
	fmt.Println("SECTION 2: Pointers and Functions")
	// Passing a pointer to a function.
	num := 10
	fmt.Println("Before increment:", num)
	increment(&num)
	fmt.Println("After increment:", num)

	// Returning a pointer from a function.
	newPtr := createPointer(20)
	fmt.Printf("Returned pointer value: %d\n", *newPtr)
	fmt.Println()

	// SECTION 3: Pointers and Structs
	fmt.Println("SECTION 3: Pointers and Structs")
	type Person struct {
		name string
		age  int
	}

	// Struct pointer.
	person := Person{name: "Alice", age: 25}
	personPtr := &person
	fmt.Printf("Original Struct: %+v\n", person)
	personPtr.age = 30 // Modify through pointer.
	fmt.Printf("Updated Struct: %+v\n", person)
	fmt.Println()

	// SECTION 4: Pointer to Pointer
	fmt.Println("SECTION 4: Pointer to Pointer")
	a := 5
	p1 := &a  // Pointer to a
	p2 := &p1 // Pointer to pointer
	fmt.Printf("Value of a: %d, Address of a: %p\n", a, &a)
	fmt.Printf("Value of p1: %p, Value at p1: %d\n", p1, *p1)
	fmt.Printf("Value of p2: %p, Value at p2: %p, Value at *p2: %d\n", p2, *p2, **p2)
	fmt.Println()

	// SECTION 5: Nil Pointers
	fmt.Println("SECTION 5: Nil Pointers")
	var nilPtr *int
	fmt.Println("Value of nilPtr:", nilPtr)
	if nilPtr == nil {
		fmt.Println("nilPtr is nil")
	}
	fmt.Println()

	// SECTION 6: Pointers and Arrays
	fmt.Println("SECTION 6: Pointers and Arrays")
	array := [3]int{10, 20, 30}
	arrayPtr := &array // Pointer to array.
	fmt.Printf("Original Array: %v\n", array)
	arrayPtr[1] = 99 // Modify through pointer.
	fmt.Printf("Updated Array: %v\n", array)
	fmt.Println()

	// SECTION 7: Advanced Pointer Concepts
	fmt.Println("SECTION 7: Advanced Pointer Concepts")
	// Pointer arithmetic (unsupported in Go, but can be mimicked with slices).
	slice := []int{1, 2, 3, 4, 5}
	ptrToSlice := &slice[0]
	fmt.Printf("First element of slice: %d, Address: %p\n", *ptrToSlice, ptrToSlice)

	// No direct pointer arithmetic, but you can access elements using indices.
	for i := range slice {
		fmt.Printf("Element %d: %d, Address: %p\n", i, slice[i], &slice[i])
	}
	fmt.Println()

	// SECTION 8: Comparison and Key Points
	fmt.Println("SECTION 8: Comparison and Key Points")
	fmt.Println("1. Pointers allow efficient modification without copying values.")
	fmt.Println("2. Use & to get the address of a variable.")
	fmt.Println("3. Use * to dereference a pointer and access the value.")
	fmt.Println("4. Nil pointers must be checked before dereferencing.")
	fmt.Println("5. Pointers are safe in Go due to the lack of pointer arithmetic.")

}

// increment modifies a value using a pointer.
func increment(num *int) {
	*num++
}

// createPointer returns a pointer to an integer.
func createPointer(value int) *int {
	return &value
}
