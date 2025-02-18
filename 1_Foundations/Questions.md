# Go Programming Practice Questions

## 1. Hello World
1. Write a program to print "Hello, Go!" to the console.
2. Modify the program to take the user's name as input and greet them with "Hello, [Name]!".
3. Write a program to print a multiline string using a single `fmt.Println` statement.

---

## 2. Variables and Constants
1. Declare variables of type `int`, `float64`, and `string`, and initialize them with values. Print the values and their types.
2. Create a constant for Pi (3.14159) and use it to calculate the circumference of a circle with a given radius.
3. Demonstrate the difference between `var` declaration and short variable declaration (`:=`).

---

## 3. Control Statements
1. Write a program to check whether a given number is even or odd.
2. Create a program to print all numbers from 1 to 100 that are divisible by 3 using a `for` loop.
3. Use a `switch` statement to print the day of the week based on a number input (1 for Monday, 2 for Tuesday, etc.).
4. Write a program to calculate the factorial of a number using a loop.

---

## 4. Functions
1. Write a function that takes two integers and returns their sum, difference, product, and quotient.
2. Create a recursive function to calculate the Fibonacci sequence up to the nth term.
3. Write a function that swaps two numbers using pointers.
4. Define and use an anonymous function to calculate the square of a number.

---

## 5. Arrays, Slices, and Maps
1. Create an array of 5 integers and calculate their sum.
2. Write a program to append elements to a slice dynamically and print its length and capacity after each addition.
3. Create a map to store student names as keys and their grades as values. Add, retrieve, and delete entries from the map.
4. Write a program to find the largest and smallest numbers in a slice.

---

## 6. Structs and Methods
1. Define a struct `Person` with fields `Name`, `Age`, and `City`. Create an instance and print its details.
2. Add a method to the `Person` struct to print a greeting message like "Hi, I'm [Name] from [City]".
3. Create a struct `Rectangle` with fields `Length` and `Width`. Add a method to calculate and return its area.
4. Implement a pointer receiver method to update the `Age` of a `Person` struct.

---

## 7. Pointers
1. Write a program to demonstrate how a pointer can modify the value of a variable.
2. Create a function that increments a number using a pointer.
3. Define a pointer to a struct and modify its fields using the pointer.
4. Demonstrate a pointer to a pointer in Go and print all dereferenced values.

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

---

### Additional Notes
- Focus on proper formatting and naming conventions.
- Experiment with error handling where applicable.
- Use comments to explain your code.
