# Functions in Go

This module covers the use of functions in Go, including basic functions, multiple return values, variadic functions, anonymous functions, closures, recursion, function types, and higher-order functions. Mastering these concepts is essential for modular and reusable Go code.

---

## 1. What is a Function?
A function is a reusable block of code that performs a specific task. Functions help break programs into smaller, manageable pieces, improving readability, maintainability, and testability.

### Why Functions Matter
- Promote code reuse and modularity.
- Make code easier to test and debug.
- Enable abstraction and separation of concerns.

### Real-World Analogy
Think of a function as a kitchen appliance: you give it ingredients (inputs), it performs a task (process), and gives you a result (output).

---

## 2. Basic Functions
Functions in Go are defined using the `func` keyword. Go requires explicit parameter and return types, making code clear and type-safe.

```go
func greet(name string) string {
    return "Hello, " + name
}
```

- **Why explicit types?**
  - Prevents ambiguity and bugs.
  - Makes function contracts clear.

---

## 3. Multiple Return Values
Go functions can return multiple values, a feature that is rare in many languages.

```go
func calculate(a, b int) (int, int) {
    return a + b, a - b
}
```

- **Why multiple returns?**
  - Enables returning both result and error (idiomatic in Go).
  - Reduces the need for custom structs or out-parameters.

**Example:**
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

---

## 4. Variadic Functions
Variadic functions accept a variable number of arguments.

```go
func sumAll(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

- **Why variadic?**
  - Flexible APIs (e.g., `fmt.Println`)
  - Avoids the need for slice arguments in simple cases.

---

## 5. Anonymous Functions
Anonymous functions (lambdas) are functions without a name. They are useful for short-lived operations, callbacks, or closures.

```go
square := func(x int) int { return x * x }
fmt.Println(square(5))
```

---

## 6. Closures
A closure is a function that captures variables from its surrounding scope.

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

- **Why closures?**
  - Enable stateful functions.
  - Useful for event handlers, iterators, and more.

---

## 7. Recursion
A recursive function is one that calls itself. Recursion is useful for problems that can be broken down into similar subproblems (e.g., factorial, tree traversal).

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

- **Best practices:**
  - Always define a base case to avoid infinite recursion.
  - Be mindful of stack overflows for deep recursion.

---

## 8. Function Types and Higher-Order Functions
Go allows you to define types for functions and pass them as arguments (higher-order functions).

```go
type operation func(int, int) int

func apply(op operation, a, b int) int {
    return op(a, b)
}
```

- **Why higher-order functions?**
  - Enable flexible APIs (e.g., sorting, filtering, mapping).
  - Promote code reuse and abstraction.

---

## 9. Best Practices
- Keep functions small and focused on a single task.
- Use descriptive names for functions and parameters.
- Document parameters, return values, and side effects.
- Prefer returning errors as the last return value.
- Use recursion judiciously; prefer iteration for most cases in Go.

---

## Go vs. Other Languages
- Go does not support default arguments or function overloading (unlike Python/C++). This keeps APIs simple and explicit.
- Go’s error handling is explicit, using multiple return values instead of exceptions.
- First-class functions and closures are similar to JavaScript and Python, but with static typing.

---

## Practice & Conceptual Questions
1. Call `greet()` with your own name.
2. Create a new function that returns product and quotient of two numbers.
3. Modify `sumAll` to return the average as well.
4. Write an anonymous function that returns the square of a number.
5. Create a closure that accumulates a sum.
6. Write a recursive function to compute Fibonacci numbers.
7. Create a new operation type function that divides two numbers.
8. Use a higher-order function with subtraction logic.
9. What are the advantages of using closures in Go?
10. How does Go handle multiple return values? Give a practical example.

---

## Further Reading
- [Go by Example: Functions](https://gobyexample.com/functions)
- [Effective Go: Functions](https://golang.org/doc/effective_go#functions) 