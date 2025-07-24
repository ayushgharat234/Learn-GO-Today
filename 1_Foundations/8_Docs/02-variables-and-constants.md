# Variables and Constants in Go

This module covers the foundational concepts of variables, constants, type inference, default values, and data types in Go. Understanding these is crucial for writing idiomatic and robust Go code.

---

## 1. Variables
Variables are named storage locations for values. In Go, variables are statically typed, meaning their type is known at compile time. This helps catch errors early and improves code reliability.

### Explicit Type Declaration
```go
var name string = "Alice"
var age int = 30
var height float64 = 5.9
var isStudent bool = false
```
- **Why explicit types?**
  - Improves code clarity, especially in large codebases.
  - Useful when the type is not obvious from the value.
  - At package scope, only `var` is allowed.

### Type Inference
```go
name := "Bob"
age := 25
height := 6.2
isStudent := true
```
- **Why type inference?**
  - Reduces verbosity for local variables.
  - Makes code concise while retaining type safety.
  - Go infers the type from the right-hand side value.

### Zero Values
Uninitialized variables in Go have default "zero values":
- `string`: ""
- `int`: 0
- `float64`: 0.0
- `bool`: false

**Why zero values?**
- Prevents undefined behavior (unlike C/C++ where uninitialized variables can have garbage values).
- Makes code safer and more predictable.

### Variable Scope
- Variables declared inside a function are local to that function.
- Variables declared at the package level are accessible throughout the package.

---

## 2. Constants
Constants are immutable values known at compile time. They are declared with the `const` keyword and cannot be changed after initialization.

```go
const pi float64 = 3.14159
const gravity = 9.8
```

### Why use constants?
- Prevent accidental modification of values that should not change (e.g., mathematical constants, configuration values).
- Improve code readability and maintainability.
- Enable compiler optimizations.

### iota: Enumerated Constants
`iota` is a Go keyword for creating sequentially incremented constants, often used for enums.
```go
const (
  _ = iota
  Low
  Medium
  High
)
```
- **Why iota?**
  - Simplifies the creation of related constant values.
  - Reduces manual errors in assigning values.

---

## 3. Data Types
Go provides a rich set of built-in types:
- **Integers:** `int`, `int8`, `int16`, `int32`, `int64`, `uint`, etc.
- **Floating Point:** `float32`, `float64`
- **Complex:** `complex64`, `complex128`
- **Rune:** Alias for `int32`, represents a Unicode code point.
- **Byte:** Alias for `uint8`.

**Why so many types?**
- Allows for memory-efficient and precise data representation.
- Useful for systems programming, networking, and performance-critical code.

---

## 4. Type Conversion
Go requires explicit type conversion:
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
- **Why explicit conversion?**
  - Prevents silent bugs due to implicit type coercion (common in dynamic languages).
  - Makes code intentions clear.

---

## 5. String Operations
- Concatenation: `greeting + ", " + audience + "!"`
- Raw strings: Use backticks for multi-line or unescaped strings.

**Go and Unicode:**
- Strings in Go are UTF-8 encoded by default.
- Use `rune` for Unicode code points.

---

## 6. Best Practices
- Use `var` for zero values or explicit declarations.
- Use `const` for values that never change.
- Always convert types explicitly.
- Understand zero values—they are idiomatic in Go.
- Prefer short variable names for short scopes, descriptive names for longer scopes.

---

## Real-World Analogy
Think of variables as labeled boxes in a warehouse. Each box (variable) can only hold a specific type of item (type). Constants are like boxes with a lock—once you put something in, it can never be changed.

---

## Practice & Conceptual Questions
1. Declare a variable for a bank balance as `float64` with value `1000.75`.
2. Create a boolean variable to store if a server is active.
3. Define a string variable for city name.
4. Use `:=` to define a temperature value in Celsius.
5. Create a variable to store whether a file is read-only.
6. Declare a variable for a default score and print it.
7. What is the default value of a declared but uninitialized boolean?
8. Define a constant for a timeout value of 5 seconds.
9. Use `iota` to define log levels: DEBUG, INFO, WARN, ERROR.
10. Declare a complex number and print its real and imaginary parts.
11. Convert a Unicode string to runes and print the count.
12. Convert a `float64` to `int` and print both.
13. Why is explicit conversion required in Go?
14. Create a boolean to determine if a user can access premium features (age > 21).
15. Combine two strings to form a message.
16. Print a multi-line log message using raw strings.
17. What are zero values in Go? Give examples for different types.
18. Why does Go require explicit type conversion? Show an example where implicit conversion would cause a bug.
19. What is the difference between a constant and a variable in Go? When should you use each?

---

## Further Reading
- [Go Tour: Variables](https://go.dev/tour/basics/8)
- [Go Tour: Constants](https://go.dev/tour/basics/15)
- [Go Blog: Variable Declaration Syntax](https://go.dev/blog/declaration-syntax)
- [Go Blog: Go and Unicode Strings](https://go.dev/blog/strings)
- [Go by Example: Variables](https://gobyexample.com/variables)
- [Go by Example: Constants](https://gobyexample.com/constants)
- [Go by Example: Type Conversions](https://go.dev/tour/basics/13) 