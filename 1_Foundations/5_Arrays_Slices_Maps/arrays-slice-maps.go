// Package main demonstrates arrays, slices, and maps in Go,
// along with their key operations, properties, and use cases.
package main

import "fmt"

func main() {
	// SECTION 1: Arrays
	fmt.Println("SECTION 1: Arrays")
	// Arrays have a fixed size and store elements of the same type.
	var arr [5]int // Declare an array of integers with a fixed size of 5.
	arr[0] = 10    // Assign a value to the first index.
	arr[1] = 20    // Assign a value to the second index.

	fmt.Println("Array:", arr)
	fmt.Printf("Length of array: %d\n", len(arr)) // Get the length of the array.

	// Declare and initialize an array.
	initializedArray := [3]int{1, 2, 3}
	fmt.Println("Initialized Array:", initializedArray)

	// Loop through an array.
	for i, value := range initializedArray {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}
	fmt.Println()

	// SECTION 2: Slices
	fmt.Println("SECTION 2: Slices")
	// Slices are dynamic and can grow or shrink. They are built on top of arrays.

	slice := []int{10, 20, 30} // Declare and initialize a slice.
	fmt.Println("Slice:", slice)
	fmt.Printf("Length of slice: %d, Capacity of slice: %d\n", len(slice), cap(slice))

	// Append elements to a slice.
	slice = append(slice, 40, 50)
	fmt.Println("After appending elements:", slice)

	// Create a slice from an array.
	array := [5]int{1, 2, 3, 4, 5}
	sliceFromArray := array[1:4] // Create a slice of elements from index 1 to 3.
	fmt.Println("Slice from array:", sliceFromArray)

	// Modifying a slice modifies the underlying array.
	sliceFromArray[0] = 99
	fmt.Println("Modified Slice:", sliceFromArray)
	fmt.Println("Underlying Array:", array)

	// Copy slices.
	newSlice := make([]int, len(slice))
	copy(newSlice, slice) // Copy contents of one slice to another.
	fmt.Println("Copied Slice:", newSlice)
	fmt.Println()

	// SECTION 3: Maps
	fmt.Println("SECTION 3: Maps")
	// Maps are key-value pairs.

	myMap := make(map[string]int) // Declare and initialize an empty map.
	myMap["Alice"] = 25           // Add a key-value pair.
	myMap["Bob"] = 30             // Add another key-value pair.
	fmt.Println("Map:", myMap)

	// Access a value by its key.
	fmt.Printf("Age of Alice: %d\n", myMap["Alice"])

	// Check if a key exists.
	value, exists := myMap["Charlie"]
	if exists {
		fmt.Printf("Age of Charlie: %d\n", value)
	} else {
		fmt.Println("Key 'Charlie' does not exist in the map.")
	}

	// Delete a key-value pair.
	delete(myMap, "Bob")
	fmt.Println("Map after deleting 'Bob':", myMap)

	// Iterate over a map.
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Nested maps.
	nestedMap := map[string]map[string]int{
		"GroupA": {"Alice": 25, "Bob": 30},
		"GroupB": {"Charlie": 35},
	}
	fmt.Println("Nested Map:", nestedMap)
	fmt.Println()

	// SECTION 4: Advanced Slice Operations
	fmt.Println("SECTION 4: Advanced Slice Operations")
	// Reslicing
	advancedSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice:", advancedSlice)
	fmt.Println("Resliced (1:3):", advancedSlice[1:3])
	fmt.Println("Resliced (2:):", advancedSlice[2:])
	fmt.Println("Resliced (:3):", advancedSlice[:3])

	// Appending beyond capacity
	extendedSlice := make([]int, 2, 3)
	extendedSlice[0] = 10
	extendedSlice[1] = 20
	fmt.Printf("Before appending beyond capacity: %v (len: %d, cap: %d)\n", extendedSlice, len(extendedSlice), cap(extendedSlice))
	extendedSlice = append(extendedSlice, 30, 40) // Appends beyond initial capacity.
	fmt.Printf("After appending beyond capacity: %v (len: %d, cap: %d)\n", extendedSlice, len(extendedSlice), cap(extendedSlice))
	fmt.Println()

	// SECTION 5: Comparison Between Arrays, Slices, and Maps
	fmt.Println("SECTION 5: Comparison")
	fmt.Println("1. Arrays are fixed in size, while slices are dynamic.")
	fmt.Println("2. Arrays cannot be resized, but slices can grow/shrink.")
	fmt.Println("3. Maps are unordered collections of key-value pairs, while arrays/slices are ordered.")
}
