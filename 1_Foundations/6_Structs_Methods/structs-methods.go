// Package main demonstrates the concepts of structs and methods in Go.
// It covers struct definition, initialization, embedding, methods, and advanced usage.
package main

import "fmt"

// Person defines a structure with fields Name and Age.
type Person struct {
	Name string // Name of the person
	Age  int    // Age of the person
}

// Greet is a method that belongs to the Person struct.
// It uses the receiver (p Person) to access struct fields and returns a greeting.
func (p Person) Greet() string {
	return "Hi, I'm " + p.Name
}

// UpdateAge updates the age of the person.
// This uses a pointer receiver to modify the original struct.
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

// Employee is a struct that embeds Person and adds additional fields.
type Employee struct {
	Person             // Embedding the Person struct
	Position   string  // Job position of the employee
	Salary     float64 // Salary of the employee
	Department string  // Department of the employee
}

// DisplayDetails is a method of Employee that prints detailed information.
func (e Employee) DisplayDetails() {
	fmt.Printf("Name: %s\nAge: %d\nPosition: %s\nSalary: %.2f\nDepartment: %s\n",
		e.Name, e.Age, e.Position, e.Salary, e.Department)
}

// Company defines a struct with nested fields.
type Company struct {
	Name      string
	Employees []Employee // Slice of Employee structs
}

// AddEmployee adds a new employee to the company.
func (c *Company) AddEmployee(e Employee) {
	c.Employees = append(c.Employees, e)
}

func main() {
	// SECTION 1: Basic Struct Usage
	fmt.Println("SECTION 1: Basic Struct Usage")
	person := Person{Name: "Alice", Age: 25} // Initializing a struct with field names
	fmt.Println("Person Name:", person.Name)
	fmt.Println("Person Age:", person.Age)
	fmt.Println("Greeting:", person.Greet()) // Call a method on the struct
	fmt.Println()

	// SECTION 2: Pointer Receiver and Modifying Structs
	fmt.Println("SECTION 2: Pointer Receiver and Modifying Structs")
	person.UpdateAge(30) // Using a pointer receiver to modify the struct
	fmt.Println("Updated Age:", person.Age)
	fmt.Println()

	// SECTION 3: Struct Embedding
	fmt.Println("SECTION 3: Struct Embedding")
	employee := Employee{
		Person:     Person{Name: "Bob", Age: 35},
		Position:   "Software Engineer",
		Salary:     75000.50,
		Department: "IT",
	}
	fmt.Println("Employee Details:")
	employee.DisplayDetails()
	fmt.Println()

	// SECTION 4: Anonymous Structs
	fmt.Println("SECTION 4: Anonymous Structs")
	anonymous := struct {
		Name  string
		Email string
	}{
		Name:  "Charlie",
		Email: "charlie@example.com",
	}
	fmt.Printf("Anonymous Struct - Name: %s, Email: %s\n", anonymous.Name, anonymous.Email)
	fmt.Println()

	// SECTION 5: Nested Structs and Composition
	fmt.Println("SECTION 5: Nested Structs and Composition")
	company := Company{Name: "Tech Corp"}
	company.AddEmployee(employee)
	company.AddEmployee(Employee{
		Person:     Person{Name: "Eve", Age: 29},
		Position:   "Data Scientist",
		Salary:     90000.00,
		Department: "Analytics",
	})

	fmt.Printf("Company: %s\n", company.Name)
	fmt.Println("Employees:")
	for _, emp := range company.Employees {
		emp.DisplayDetails()
		fmt.Println()
	}

	// SECTION 6: Comparison of Structs
	fmt.Println("SECTION 6: Comparison of Structs")
	person1 := Person{Name: "John", Age: 40}
	person2 := Person{Name: "John", Age: 40}
	person3 := Person{Name: "Jane", Age: 40}
	fmt.Println("person1 == person2:", person1 == person2) // True if all fields match
	fmt.Println("person1 == person3:", person1 == person3) // False due to different Name
	fmt.Println()

	// SECTION 7: Advanced Struct Concepts
	fmt.Println("SECTION 7: Advanced Struct Concepts")
	// Zero Value of a Struct
	var defaultPerson Person
	fmt.Printf("Default Person: Name: %q, Age: %d\n", defaultPerson.Name, defaultPerson.Age)

	// Creating a pointer to a struct
	personPointer := &Person{Name: "Diana", Age: 22}
	fmt.Printf("Pointer to Struct - Name: %s, Age: %d\n", personPointer.Name, personPointer.Age)
}
