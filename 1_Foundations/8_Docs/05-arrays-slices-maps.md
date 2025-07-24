# Arrays, Slices, and Maps in Go

This module covers the three primary data structures for collections in Go: arrays, slices, and maps. Understanding their properties and use cases is essential for effective Go programming.

---

## 1. Arrays
Arrays are fixed-size, ordered collections of elements of the same type. They are value types, meaning assignments copy the entire array.

### Syntax
```go
var arr [5]int
arr[0] = 10
arr[1] = 20
```

### Why Arrays?
- Useful for low-level programming, memory control, and when the size is known at compile time.
- Arrays are rarely used directly in Go application code, but are the foundation for slices.

### Real-World Analogy
Think of an array as a row of lockers: each locker (element) has a fixed position and can only store one item of a specific type.

---

## 2. Slices
Slices are dynamic, flexible views into arrays. They are the most commonly used collection type in Go.

### Syntax
```go
slice := []int{10, 20, 30}
slice = append(slice, 40, 50)
```

### Why Slices?
- Dynamic size: can grow or shrink as needed.
- Reference type: passing a slice to a function does not copy the underlying data.
- Backed by arrays, but provide much more flexibility.

### Capacity and Length
- `len(slice)`: number of elements.
- `cap(slice)`: capacity of the underlying array.
- Slices can be resliced, appended, and copied.

### Memory and Performance
- Appending beyond capacity allocates a new underlying array.
- Slices are efficient for most use cases, but be mindful of memory leaks if large arrays are referenced by small slices.

### Real-World Analogy
A slice is like a window into a row of lockers (an array): you can move the window, resize it, and see only part of the row.

---

## 3. Maps
Maps are unordered collections of key-value pairs. They are reference types and are implemented as hash tables.

### Syntax
```go
myMap := make(map[string]int)
myMap["Alice"] = 25
myMap["Bob"] = 30
```

### Why Maps?
- Fast lookups, inserts, and deletions by key.
- Ideal for dictionaries, caches, and associative arrays.
- Keys must be comparable (e.g., strings, numbers, but not slices or maps).

### Real-World Analogy
A map is like a dictionary: you look up a word (key) to find its definition (value).

---

## 4. Advanced Slice Operations
- Reslicing: `slice[1:3]`, `slice[2:]`, `slice[:3]`
- Appending beyond capacity creates a new underlying array.
- Use `make([]int, length, capacity)` for custom capacity.

---

## 5. Comparison
- **Arrays:** fixed size, value type, ordered.
- **Slices:** dynamic, reference type, ordered.
- **Maps:** dynamic, reference type, unordered.

### Choosing Between Them
- Use arrays for fixed-size, low-level data.
- Use slices for most dynamic collections.
- Use maps for fast lookups and key-value storage.

---

## 6. Best Practices
- Avoid using arrays unless you need fixed size and value semantics.
- Use slices for most collections.
- Use maps for associative data.
- Always check if a key exists in a map before using its value.
- Be careful with slices referencing large arrays (can cause memory leaks).

---

## Go vs. Other Languages
- Go’s slices are similar to Python lists or JavaScript arrays, but with explicit capacity and underlying array semantics.
- Maps are similar to dictionaries in Python or objects in JavaScript, but with stricter key requirements.

---

## Practice & Conceptual Questions
1. Create an array of strings and print each character in reverse order.
2. Use append and copy to merge two slices.
3. Create a map of countries with nested maps of cities and populations.
4. Create a slice, append till it doubles its capacity and print it at each step.
5. Summarize pros and cons of each data structure in your own words.
6. What are the differences between value types and reference types in Go? How does this affect arrays, slices, and maps?

---

## Further Reading
- [Go by Example: Arrays](https://gobyexample.com/arrays)
- [Go by Example: Slices](https://gobyexample.com/slices)
- [Go by Example: Maps](https://gobyexample.com/maps) 