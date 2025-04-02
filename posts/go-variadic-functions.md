Title: Go Variadic Functions
Date: 02-Apr-2025

Variadic functions in Go provide a flexible way to work with a variable number of arguments.  Similar to the concept of "rest parameters" in JavaScript, Go's variadic functions allow you to pass zero or more arguments of a specific type to a function.

**Declaration**

A variadic function is declared using an ellipsis (`...`) preceding the type name of the variadic parameter. This parameter effectively becomes a slice within the function's scope.

```go
package main

import "fmt"

func printStrings(strs ...string) {
    for _, str := range strs {
        fmt.Println(str)
    }
}

func main() {
    printStrings("hello", "world", "go")
    printStrings() // Calling with no arguments is also valid
}

```

**Usage**

The `strs` parameter inside `printStrings` is treated as a slice of strings.  You can iterate over it, check its length, or access individual elements just like any other slice.

**Passing Slices to Variadic Functions**

You can pass an existing slice to a variadic function by appending the ellipsis to the slice variable during the function call. This effectively "unpacks" the slice into individual arguments.

```go
package main

import "fmt"

func sumNumbers(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    sum := sumNumbers(numbers...) // Unpack the slice
    fmt.Println("Sum:", sum) // Output: Sum: 15


	moreNumbers := []int{10,11}
	combinedSlice := append(numbers, moreNumbers...) // append moreNumbers to numbers
	sum2 := sumNumbers(combinedSlice...) // Unpack the combined slice
	fmt.Println(sum2) // Output: 66
}

```

**Combining with Regular Parameters**

Variadic parameters can be combined with regular parameters, but the variadic parameter must always be the last parameter in the function signature.

```go
package main

import "fmt"

func printFormatted(prefix string, values ...interface{}) {
    for _, value := range values {
        fmt.Println(prefix, value)
    }
}

func main() {
    printFormatted("Value:", 1, "hello", 3.14)
}
```

In this example, `prefix` is a regular string parameter, and `values` is a variadic parameter of type `interface{}`, which allows values of any type to be passed.


Variadic functions provide a concise and elegant way to handle functions that need to accept a variable number of arguments, making your Go code more flexible and readable.
