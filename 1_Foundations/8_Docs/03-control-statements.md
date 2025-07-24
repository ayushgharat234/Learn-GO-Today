# Control Statements in Go

This module covers the essential control flow constructs in Go: conditional statements, loops, and switch-case. Mastery of these is fundamental for writing logical and efficient Go programs.

---

## 1. Conditional Statements (if-else)
Conditional statements allow your program to make decisions based on conditions. In Go, the `if` statement is simple and does not require parentheses around the condition.

### Syntax
```go
if condition {
    // code
} else if anotherCondition {
    // code
} else {
    // code
}
```

### Why Go's if-else?
- No parentheses required, making code less cluttered.
- You can declare variables in the if statement, scoped only to that block:
  ```go
  if x := compute(); x > 10 {
      // use x
  }
  ```
- Encourages short, readable blocks.

### Real-World Analogy
Think of `if-else` as a series of road signs: "If the road is closed, take the detour; else if there's traffic, take the highway; else, continue straight."

---

## 2. Loops
Go has only one looping construct: the `for` loop. It is versatile and can be used as a traditional for, a while, or an infinite loop.

### Classic for
```go
for i := 0; i < 5; i++ {
    // ...
}
```

### While-style
```go
for condition {
    // ...
}
```

### Infinite loop
```go
for {
    // ...
    if done {
        break
    }
}
```

### Why only for?
- Simplifies the language and reduces confusion.
- All loop patterns are possible with just `for`.

### Best Practices
- Use `break` to exit a loop early.
- Use `continue` to skip to the next iteration.
- Avoid deeply nested loops for readability.

---

## 3. Switch Statement
The `switch` statement in Go is a powerful alternative to long if-else chains. It can be used with or without a condition.

### Syntax
```go
switch variable {
case value1:
    // ...
case value2:
    // ...
default:
    // ...
}
```

### Unique Features
- Cases do not "fall through" by default (unlike C/C++/Java). Use `fallthrough` explicitly if needed.
- Can be used without a condition for more complex logic.
- Multiple values can be matched in a single case.

### Real-World Analogy
A switch is like a train station: depending on the track (case), the train (program) goes in a different direction.

---

## 4. Best Practices
- Prefer `switch` over long if-else chains for clarity.
- Use short variable declarations in if/switch when possible.
- Avoid infinite loops unless necessary (e.g., servers).
- Keep control flow blocks short and focused.

---

## Why Go's Control Flow?
- Go’s minimalism means fewer ways to do the same thing, making codebases more uniform.
- The lack of exceptions (no try/catch) means errors are handled explicitly, usually with return values.
- Encourages clear, linear code that is easy to follow and maintain.

---

## Practice & Conceptual Questions
1. Write a program to classify a grade into A/B/C/Fail using if-else.
2. Check if a number is positive, negative, or zero.
3. Print the factorial of a number using a loop.
4. Print the Fibonacci series up to N terms.
5. Sum even numbers from 1 to 50 using continue.
6. Create a switch case to identify days from 1–7.
7. Use fallthrough to group weather types: Sunny, Cloudy, Rainy.
8. Replace if-else ladder for marks grading with switch.
9. What are the benefits of using switch over multiple if-else statements?

---

## Further Reading
- [Go by Example: If/Else](https://gobyexample.com/if-else)
- [Go by Example: For](https://gobyexample.com/for)
- [Go by Example: Switch](https://gobyexample.com/switch) 