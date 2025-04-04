Title: Go Control Flow
Date: 04-Apr-2025

Go, like most programming languages, uses control flow statements to dictate the order in which code is executed.  These structures allow you to create branching logic, loops, and conditional execution, making your programs dynamic and responsive.

**Conditional Statements: if, else if, else**

Go's `if`, `else if`, and `else` statements work similarly to other C-style languages.  They allow you to execute blocks of code based on a boolean condition.

```go
package main

import "fmt"

func main() {
	age := 25

	if age < 18 {
		fmt.Println("Minor")
	} else if age >= 18 && age < 65 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior")
	} // Output: Adult
}

```

Parentheses around the condition are optional in Go.  The `else if` and `else` blocks are also optional.


**Switch Statements**

Go's `switch` statement provides a concise way to handle multiple conditional branches based on the value of an expression.

```go
package main

import "fmt"

func main() {
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Tuesday", "Wednesday", "Thursday": // Multiple cases can be combined
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("Almost weekend!")
	default: // Default case handles all other values
		fmt.Println("Weekend")
	} // Output: Midweek
}

```

Go's `switch` has implicit breaks, meaning execution doesn't fall through to the next case.  You can explicitly use `fallthrough` if you need this behavior.


**For Loops**

The `for` loop is Go's primary looping construct.  It's incredibly versatile and can be used in various ways.

```go
package main

import "fmt"

func main() {

	// Classic for loop with initializer, condition, and post statement
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Condition-only loop (like a while loop)
	j := 0
	for j < 3 {
		fmt.Println(j)
		j++
	}

	// Infinite loop (use break to exit)
	for {
		fmt.Println("Looping...")
		break // Exit the loop
	}

	// Ranging over an array/slice
	arr := []string{"apple", "banana", "cherry"}
	for index, value := range arr {
		fmt.Println(index, value)
	}
}
```


**Break and Continue**

The `break` statement is used to exit a loop prematurely. The `continue` statement skips the current iteration and proceeds to the next.

These control flow statements are fundamental building blocks for writing structured and efficient Go programs. By mastering these constructs, you can create complex logic and control the execution path of your code effectively.
