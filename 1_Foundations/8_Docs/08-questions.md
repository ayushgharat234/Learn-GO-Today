# Go Programming Practice & Conceptual Questions

This section provides a comprehensive set of practice and conceptual questions for each foundational topic in Go. These are designed to deepen your understanding and prepare you for real-world, cloud-native Go development.

---

## 1. Hello World
1. Write a program to print "Hello, Go!" to the console.
2. Modify the program to take the user's name as input and greet them with "Hello, [Name]!".
3. Write a program to print a multiline string using a single `fmt.Println` statement.
4. What is the purpose of the `main` package and `main()` function in Go?
5. How does Go handle Unicode and UTF-8 in string literals?

---

## 2. Variables and Constants
1. Declare variables of type `int`, `float64`, and `string`, and initialize them with values. Print the values and their types.
2. Create a constant for Pi (3.14159) and use it to calculate the circumference of a circle with a given radius.
3. Demonstrate the difference between `var` declaration and short variable declaration (`:=`).
4. What are zero values in Go? Give examples for different types.
5. Use `iota` to define an enum for log levels: DEBUG, INFO, WARN, ERROR.
6. Why does Go require explicit type conversion? Show an example where implicit conversion would cause a bug.
7. Write a function that takes a string and returns its rune count and Unicode code points.
8. What is the difference between a constant and a variable in Go? When should you use each?

---

## 3. Control Statements
1. Write a program to check whether a given number is even or odd.
2. Create a program to print all numbers from 1 to 100 that are divisible by 3 using a `for` loop.
3. Use a `switch` statement to print the day of the week based on a number input (1 for Monday, 2 for Tuesday, etc.).
4. Write a program to calculate the factorial of a number using a loop.
5. Write a program to classify a grade into A/B/C/Fail using if-else.
6. Check if a number is positive, negative, or zero.
7. Print the Fibonacci series up to N terms using a loop.
8. Sum even numbers from 1 to 50 using continue.
9. Use fallthrough in a switch to group weather types: Sunny, Cloudy, Rainy.
10. Replace an if-else ladder for marks grading with a switch statement.
11. What are the benefits of using switch over multiple if-else statements?

---

## 4. Functions
1. Write a function that takes two integers and returns their sum, difference, product, and quotient.
2. Create a recursive function to calculate the Fibonacci sequence up to the nth term.
3. Write a function that swaps two numbers using pointers.
4. Define and use an anonymous function to calculate the square of a number.
5. Modify a variadic function to return both the sum and average of its arguments.
6. Write a closure that accumulates a running total.
7. Create a function type and use it to pass different operations (add, subtract, multiply, divide) to a higher-order function.
8. What are the advantages of using closures in Go?
9. How does Go handle multiple return values? Give a practical example.

---

## 5. Arrays, Slices, and Maps
1. Create an array of 5 integers and calculate their sum.
2. Write a program to append elements to a slice dynamically and print its length and capacity after each addition.
3. Create a map to store student names as keys and their grades as values. Add, retrieve, and delete entries from the map.
4. Write a program to find the largest and smallest numbers in a slice.
5. Create an array of strings and print each character in reverse order.
6. Use append and copy to merge two slices.
7. Create a map of countries with nested maps of cities and populations.
8. Create a slice, append till it doubles its capacity and print it at each step.
9. Summarize pros and cons of arrays, slices, and maps in your own words.
10. What are the differences between value types and reference types in Go? How does this affect arrays, slices, and maps?

---

## 6. Structs and Methods
1. Define a struct `Person` with fields `Name`, `Age`, and `City`. Create an instance and print its details.
2. Add a method to the `Person` struct to print a greeting message like "Hi, I'm [Name] from [City]".
3. Create a struct `Rectangle` with fields `Length` and `Width`. Add a method to calculate and return its area.
4. Implement a pointer receiver method to update the `Age` of a `Person` struct.
5. Create a new struct called `Car` with fields Make, Model, and Year.
6. Add a method to `Car` called `Start` that prints a message using the car's Make and Model.
7. Create a new struct called `Garage` that contains a slice of Cars and a method to add new cars.
8. Modify the `Person` struct to include an Email field and update the Greet method to use it.
9. Compare two Employee structs and explain when two structs are considered equal in Go.
10. What is struct embedding and how does it differ from inheritance in other languages?

---

## 7. Pointers
1. Write a program to demonstrate how a pointer can modify the value of a variable.
2. Create a function that increments a number using a pointer.
3. Define a pointer to a struct and modify its fields using the pointer.
4. Demonstrate a pointer to a pointer in Go and print all dereferenced values.
5. Write a function `swap(a, b *int)` that swaps two integers using pointers.
6. Create a struct `Rectangle` with Width and Height. Write a method using pointer receiver to scale its size.
7. Declare a nil pointer to a float64 and safely assign it a value inside a function.
8. Use a pointer to update the second element in a slice of strings.
9. Declare a variable of type pointer to pointer (`**int`) and print the dereferenced value.
10. (Advanced) Create a function that takes a slice and modifies it using a pointer to the first element.
11. Why does Go disallow pointer arithmetic? What are the safety benefits?

---

## General Challenges
1. Combine all concepts to create a mini project:
   - Define a `Student` struct with fields `Name`, `Subjects`, and `Grades` (map of subject to grade).
   - Write functions to:
     - Add a subject and grade.
     - Calculate the average grade.
     - Print the student's details in a formatted way.
2. Write a program to manage an inventory system:
   - Use a map to store item names as keys and their quantities as values.
   - Add functions to add items, update quantities, and display the inventory.
3. Create a program that simulates a basic calculator supporting addition, subtraction, multiplication, and division, implemented using functions.
4. Build a RESTful API in Go (using net/http) to manage a list of books (CRUD operations). (Bonus: Use structs, slices, and maps.)
5. Design a CLI tool in Go that reads a CSV file and summarizes the data (e.g., average, min, max for numeric columns).

---

### Additional Notes
- Focus on proper formatting and naming conventions.
- Experiment with error handling where applicable.
- Use comments to explain your code.
- Try to relate each concept to real-world, cloud-native scenarios (e.g., config management, API design, data processing).
