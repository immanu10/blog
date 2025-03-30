Title: Go Constants
Date: 30-Mar-2025

Constants in Go represent immutable values determined at compile time.  They can be of various basic types like integers, floating-point numbers, characters, strings, and booleans.  Unlike variables, constants cannot be changed after they are declared.

```go
package main

import "fmt"

const (
	Pi       = 3.14159
	Language = "Go"
	Truth    = true
)

func main() {
	fmt.Println("Value of Pi:", Pi)
	fmt.Println("Language:", Language)
	fmt.Println("Is it true?", Truth)

	// The following line would result in a compile-time error
	// because you cannot modify a constant:
	// Pi = 3.14

	// Typed constants
	const typedInt int = 10
	fmt.Println("Typed constant:", typedInt)


	// Untyped constants
	const untypedFloat = 20.5
	var float32Var float32 = float32(untypedFloat)
	var float64Var float64 = untypedFloat
	fmt.Println("Untyped constant as float32:", float32Var)
	fmt.Println("Untyped constant as float64:", float64Var)
}

```

This code snippet demonstrates a few key aspects of constants:

* **Declaration:** Constants are declared using the `const` keyword.  You can group related constants together using parentheses.
* **Naming Convention:** Constant names typically follow the same rules as variable names, but with an uppercase first letter for exported constants (those accessible from other packages).
* **Type Inference:** Go can often infer the type of a constant based on its value, but you can explicitly specify the type if needed.  The examples include both typed (e.g., `typedInt`) and untyped constants (e.g., `untypedFloat`).
* **Immutability:** The commented-out line `Pi = 3.14` illustrates the fundamental principle that constants cannot be changed after they are declared.  Attempting to do so will lead to a compile-time error.
* **Untyped Constants:** Untyped constants take on the type of the context they are used in. This allows for flexible use in various situations, as shown with the `untypedFloat` example.

Using constants makes your code more readable, maintainable, and less prone to errors by providing named values that won't accidentally be modified during program execution. They are particularly useful for defining values that have a special meaning or are used repeatedly throughout your code.
