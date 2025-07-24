# Structs and Methods in Go

This module covers the use of structs and methods in Go, including struct definition, initialization, embedding, methods, pointer receivers, comparison, and advanced usage. These are essential for modeling real-world entities and behaviors in Go.

---

## 1. What is a Struct?
A struct is a composite data type that groups together variables (fields) under a single name. Structs are Go’s way of modeling real-world entities and are the foundation for data modeling in Go.

### Syntax
```go
type Person struct {
    Name string
    Age  int
}
```

### Why Structs?
- Enable grouping related data together.
- Provide a foundation for methods and interfaces.
- Used for configuration, data transfer, and more.

### Real-World Analogy
A struct is like a form: it has labeled fields (Name, Age) that you fill in with information.

---

## 2. Methods
Methods are functions with a receiver argument. The receiver can be a value or a pointer to a struct.

```go
func (p Person) Greet() string {
    return "Hi, I'm " + p.Name
}

func (p *Person) UpdateAge(newAge int) {
    p.Age = newAge
}
```

- **Value receiver:** operates on a copy of the struct.
- **Pointer receiver:** can modify the original struct.

---

## 3. Struct Embedding and Composition
Go does not have traditional inheritance. Instead, it uses composition via struct embedding.

```go
type Employee struct {
    Person
    Position string
    Salary   float64
}
```

- **Why embedding?**
  - Promotes code reuse without the complexity of inheritance.
  - Enables "is-a" and "has-a" relationships.

---

## 4. Anonymous Structs
Anonymous structs are useful for quick, one-off data structures, especially in tests or small scopes.

---

## 5. Nested Structs and Composition
Structs can contain other structs or slices of structs, enabling complex data models.

---

## 6. Comparison and Equality
- Structs are comparable if all their fields are comparable.
- Two structs are equal if all their fields are equal.
- Use `==` for comparison, but be careful with fields that are slices, maps, or functions (not comparable).

---

## 7. Advanced Concepts
- **Zero value:** All fields are set to their zero values when a struct is declared but not initialized.
- **Pointers to structs:** Efficient for passing large structs to functions and for methods that modify the struct.

---

## 8. Best Practices
- Use pointer receivers for methods that modify the struct or for efficiency with large structs.
- Use struct embedding for composition, not inheritance.
- Prefer explicit field names for clarity.
- Document struct fields and methods.

---

## Go vs. Other Languages
- Go does not have classes or inheritance. Instead, it uses structs, methods, and interfaces for composition and polymorphism.
- This approach encourages composition over inheritance, leading to simpler and more maintainable code.

---

## Practice & Conceptual Questions
1. Create a new struct called `Car` with fields Make, Model, and Year.
2. Add a method to `Car` called `Start` that prints a message using the car's Make and Model.
3. Create a new struct called `Garage` that contains a slice of Cars and a method to add new cars.
4. Modify the `Person` struct to include an Email field and update the Greet method to use it.
5. Compare two Employee structs and explain when two structs are considered equal in Go.
6. What is struct embedding and how does it differ from inheritance in other languages?

---

## Further Reading
- [Go by Example: Structs](https://gobyexample.com/structs)
- [Go by Example: Methods](https://gobyexample.com/methods)